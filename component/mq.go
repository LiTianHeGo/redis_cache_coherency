package component

/*
 * 抽象消息队列组件(以RabbitMQ为例)
 */
type RabbitMQServer struct {
}

func (mq *RabbitMQServer) SendMsg(key string, data []byte) {

}
