package component

/*
 * 抽象缓存组件(以Redis为例)，需要提供查询数据Get、添加数据Put等能力
 */
type RedisServer struct {
}

//查询数据
func (cache *RedisServer) Get(key string) (interface{}, error) {

	return nil, nil
}

//回写数据到cache
func (cache *RedisServer) WriteBack(key string) error {

	return nil
}

//读请求尝试回写数据到cache
func (cache *RedisServer) WriteBackByReadReq(key string) error {
	//需要结合lua脚本实现检验读请求回写机制是否可用，并根据结果决定是否执行回写
	return nil
}

//删除cache
func (cache *RedisServer) Delete(key string) error {

	return nil
}

//关闭某个key的读请求回写机制
func (cache *RedisServer) DisableWriteBack(key string) error {
	//往redis插入一条和key对应的禁用标志
	return nil
}

//恢复某个key的读请求回写机制
func (cache *RedisServer) EnableWriteBack(key string) error {
	//删除redis中和key对应的禁用标志
	return nil
}
