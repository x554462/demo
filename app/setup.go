package app

import (
	"github.com/x554462/demo/middleware/mango/library/conf"
	"github.com/x554462/go-dao/db"
)

func Setup() {
	cm := conf.MasterDatabaseConf
	cs := conf.SlaveDatabaseConf
	db.Setup(db.Conf{
		Name:     cm.Name,
		User:     cm.User,
		Password: cm.Password,
		Host:     cm.Host,
		Port:     cm.Port,
	})
	db.SetupSlave(db.Conf{
		Name:     cs.Name,
		User:     cs.User,
		Password: cs.Password,
		Host:     cs.Host,
		Port:     cs.Port,
	})
}
