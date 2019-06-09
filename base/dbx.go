package base

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"github.com/wing1000/infra"
	"github.com/wing1000/infra/logrus"
)

//dbx 数据库实例
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	Check(database)
	return database
}

//dbx数据库starter，并且设置为全局
type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	//数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	log.Info("mysql.conn url:", settings.ShortDataSourceName())
	db, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	log.Info(db.Ping())
	db.SetLogger(logrus.NewUpperLogrusLogger())
	database = db
}
