package zetcd

import (
	"github.com/Xuzan9396/zlog"
	client "github.com/soyking/e3ch"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var e3w *client.EtcdHRCHYClient

func EtchClt() *client.EtcdHRCHYClient {
	return e3w
}

func Run() {
	// initial etcd v3 client
	//e3Clt, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	e3Clt, err := clientv3.New(clientv3.Config{Endpoints: []string{viper.GetString("etcd.addr")}})
	if err != nil {
		zlog.F().Error(err)
		panic(err)
		return
	}

	// new e3ch client with namespace(rootKey)
	clt, err := client.New(e3Clt, "/e3w")
	if err != nil {
		zlog.F().Error(err)
		panic(err)
		return
	}

	// set the rootKey as directory
	err = clt.FormatRootKey()
	if err != nil {
		zlog.F().Error(err)
		panic(err)
		return
	}
	e3w = clt

	//clt.CreateDir("/dir1")
	clt.Create("/dir1/key1", "2323")
	//clt.Create("/dir", "")
	//clt.Put("/dir1/key1", "value1")
	//clt.Get("/dir1/key1")
	//clt.List("/dir1")
	//clt.Delete("/dir")
}