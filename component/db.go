package component

/*
 * 抽象DB组件(以Mysql为例)，需要提供查询数据Get()、更新数据Update()等能力
 */
type MysqlServer struct {
}

func (db *MysqlServer) Get(key string) (interface{}, error) {

	return nil, nil
}

func (db *MysqlServer) Update(key string) error {
	return nil
}

func (db *MysqlServer) Insert(key string) {

}
