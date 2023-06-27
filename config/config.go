package config

import (
	_ "embed"
	"encoding/json"
	"github.com/Xuzan9396/crawler_config/zetcd"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"runtime"
)

type Config struct {
	Name string
}

// new config
func InitConfig(name string, etcdKey ...string) error {
	c := &Config{
		Name: name,
	}
	err := c.loadConfig()
	if err != nil {
		return err
	}
	c.watchConfig()
	if len(etcdKey) > 0 {
		err = c.etcd(etcdKey[0])

	} else {
		err = c.etcd("/mysql")
	}
	if err != nil {
		return err
	}
	return nil
}

// 加载配置文件
func (c *Config) loadConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.SetConfigFile("./conf/config.yaml")
	}

	viper.SetConfigType("yaml")
	//viper.AutomaticEnv()
	//viper.SetEnvPrefix("APISERVER") // 设置前缀
	//replaces := strings.NewReplacer(".", "_")
	//viper.SetEnvKeyReplacer(replaces)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil

}
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("文件更新了", e.Name, "携程数量:", runtime.NumGoroutine())
		//log.Infof("config file changed: %s,goroutine数量：%d", e.Name, runtime.NumGoroutine())
	})
}

func (c *Config) etcd(key string) (err error) {
	err = zetcd.Run()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("etcd启动成功")
	res, err := zetcd.EtchClt().Get(key)
	if err != nil {
		log.Println(err)
		return err
	}
	var info Configs
	err = json.Unmarshal(res.Value, &info)
	if err != nil {
		log.Println(err)
		return err
	}

	viper.Set("db.username", info.Db.Username)
	viper.Set("db.password", info.Db.Password)
	viper.Set("db.addr", info.Db.Addr)
	viper.Set("db.datasename", info.Db.Datasename)

	viper.Set("db_slave.username", info.DbSlave.Username)
	viper.Set("db_slave.password", info.DbSlave.Password)
	viper.Set("db_slave.addr", info.DbSlave.Addr)
	viper.Set("db_slave.datasename", info.DbSlave.Datasename)

	viper.Set("redis.conn", info.Redis.Conn)
	viper.Set("redis.passwd", info.Redis.Passwd)
	viper.Set("redis.dbnum", info.Redis.Dbnum)

	viper.Set("mq.addr_url", info.Mq.AddrURL)
	viper.Set("mq.username", info.Mq.Username)
	viper.Set("mq.password", info.Mq.Password)
	return nil

}
