package config

type Configs struct {
	Db       Db       `json:"db"`
	DbSlave  DbSlave  `json:"db_slave"`
	Redis    Redis    `json:"redis"`
	RedisAct RedisAct `json:"redis_act"`
	Mq       Mq       `json:"mq"`
	Es       Es       `json:"es"`
	Canal    Canal    `json:"canal"`
}
type Db struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Addr       string `json:"addr"`
	Datasename string `json:"datasename"`
}

type DbSlave struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Addr       string `json:"addr"`
	Datasename string `json:"datasename"`
}
type Redis struct {
	Conn   string `json:"conn"`
	Passwd string `json:"passwd"`
	Dbnum  int    `json:"dbnum"`
}
type RedisAct struct {
	Conn   string `json:"conn"`
	Passwd string `json:"passwd"`
	Dbnum  int    `json:"dbnum"`
}
type Mq struct {
	AddrURL  string `json:"addr_url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Canal struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}
type Es struct {
	Addr               string `json:"addr"`
	Crawler_index_name string `json:"crawler_index_name"`
}
