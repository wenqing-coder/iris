package datasource

import (
	"fmt"
	"iris/conf"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var masterEngine *xorm.Engine
var slaveEngine *xorm.Engine
var groupEngine *xorm.EngineGroup
var mu sync.Mutex

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	mu.Lock()
	defer mu.Unlock()
	if masterEngine != nil {
		return masterEngine
	}
	driverSourceName := pickMysqlDriverStr(conf.MasterDbConf)
	en, err := xorm.NewEngine(conf.DriverName, driverSourceName)
	if err != nil {
		log.Fatal("dbhelper.InstanceMaster err:", err)
		return nil
	}
	en.ShowSQL(false)
	en.SetTZLocation(conf.SysTimeLocation)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	en.SetDefaultCacher(cacher)
	masterEngine = en
	return en
}

func pickMysqlDriverStr(conf conf.DbConf) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.User,
		conf.Pwd,
		conf.Host,
		conf.Port,
		conf.DbName)
}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	mu.Lock()
	defer mu.Unlock()
	if slaveEngine != nil {
		return slaveEngine
	}
	driverSourceName := pickMysqlDriverStr(conf.MasterDbConf)
	en, err := xorm.NewEngine(conf.DriverName, driverSourceName)
	if err != nil {
		log.Fatal("dbhelper.InstanceMaster err:", err)
		return nil
	}
	en.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = en
	return en
}

func InstanceDbGroup() (*xorm.EngineGroup, error) {
	if groupEngine != nil {
		return groupEngine, nil
	}
	mu.Lock()
	defer mu.Lock()
	if groupEngine != nil {
		return groupEngine, nil
	}
	driverMasterSourceName := pickMysqlDriverStr(conf.MasterDbConf)
	driverSlaveSourceName := pickMysqlDriverStr(conf.MasterDbConf)
	master, err := xorm.NewEngine(conf.DriverName, driverMasterSourceName)
	if err != nil {
		log.Println("cannot connect master mysql err ", err)
		return nil, err
	}
	slave, err := xorm.NewEngine(conf.DriverName, driverSlaveSourceName)
	if err != nil {
		log.Println("cannot connect slave mysql err ", err)
		return nil, err
	}
	slaves := []*xorm.Engine{slave}
	ge, err := xorm.NewEngineGroup(master, slaves, xorm.RandomPolicy())
	if err != nil {
		log.Println("cannot create group engine err ", err)
		return nil, err
	}
	groupEngine = ge
	return ge, nil
}
