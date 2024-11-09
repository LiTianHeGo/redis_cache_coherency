package rediscachecoherency

import (
	"context"
	"redis_cache_coherency/component"

	"golang.org/x/sync/singleflight"
)

// 缓存一致性服务模块（todo:各组件可抽象为接口）
type CacheCoherencyService struct {
	component.RedisServer    //缓存组件
	component.RabbitMQServer //消息队列组件
	component.MysqlServer    //数据库组件
}

// 处理（针对热点key的）写请求
func (svc *CacheCoherencyService) HandleWriteForHotKey(ctx context.Context, key string) error {
	//1、关闭该key的读请求回写机制
	if err := svc.RedisServer.DisableWriteBack(key); err != nil {
		return err
	}

	//2、删除缓存
	if err := svc.RedisServer.Delete(key); err != nil {
		return err
	}

	//3、执行写DB
	if err := svc.MysqlServer.Update(key); err != nil {
		svc.RedisServer.EnableWriteBack(key) //写db失败，则恢复读请求的回写机制
		return err
	}

	//4、主动回写redis
	if err := svc.RedisServer.WriteBack(key); err != nil {
		svc.RabbitMQServer.SendMsg(key, []byte("模拟数据..")) //若回写失败，结合rmq消息队列，保证最终回写缓存成功
	}

	//5、恢复该key的读请求回写机制
	return svc.RedisServer.EnableWriteBack(key)
}

// 处理（针对热点key的）读请求
func (svc *CacheCoherencyService) HandleReadForHotKey(ctx context.Context, key string) (interface{}, error) {
	//1、尝试读cache
	if v, err := svc.RedisServer.Get(key); err == nil && v != nil {
		return v, nil
	}

	//2、若读cache未命中，则结合singleflight继续读数据库
	var g singleflight.Group
	v, err, _ := g.Do(key, func() (interface{}, error) {
		if value, err := svc.MysqlServer.Get(key); err != nil {
			return nil, err
		} else {
			return value, nil
		}
	})

	if err != nil {
		return nil, err
	}
	p := v.(*Product)

	//3、尝试回写redis
	svc.RedisServer.WriteBackByReadReq(key)

	return p, nil
}
