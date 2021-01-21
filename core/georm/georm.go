package georm

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/liuyong-go/yong/core/georm/session"
)

//Engine 是 GeORM 与用户交互的入口
type Engine struct {
	db *sql.DB
}

//NewEngine 实例化
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		LogDBErr(err)
		return
	}
	if err = db.Ping(); err != nil {
		LogDBErr(err)
		return
	}
	e = &Engine{db: db}
	return
}

//Close  关闭连接
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		LogDBErr(err)
	}
}
func (engine *Engine) NewSession() *session.Session {
	return session.NewDB(engine.db)
}

//LogDBErr 打印db错误日志
func LogDBErr(err error) {
	log.Println("db err", err)
}
