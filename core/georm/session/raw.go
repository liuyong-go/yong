package session

import (
	"database/sql"
	"log"
	"strings"
)

//Session 包含数据库连接指针，sql语句和占位符对应值
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

//NewDB 实例化session
func NewDB(db *sql.DB) *Session {
	return &Session{db: db}
}

//Clear 重置sql
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

//DB 获取数据库连接指针
func (s *Session) DB() *sql.DB {
	return s.db
}

//Raw 拼写sql及绑定参数
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//Exec 执行sql
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Println("DB ERR", err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Println("DB ERR", err)
	}
	return
}
