package gomysql

import (
	"database/sql"
	"fmt"
	"gomysql/statement"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	host   string
	port   int
	user   string
	passwd string
	db     string

	conn   *sql.DB
	gormee *gorm.DB
}

func NewMysql(host string, port int, user string, passwd string, db string) (Mysql, error) {
	if conn, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, passwd, host, port, db)); err != nil {
		return Mysql{host, port, user, passwd, db, nil, nil}, err
	} else if err := conn.Ping(); err != nil {
		conn.Close()
		return Mysql{host, port, user, passwd, db, nil, nil}, err
	} else if gormee, err := gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{}); err != nil {
		conn.Close()
		return Mysql{host, port, user, passwd, db, nil, nil}, err
	} else {
		return Mysql{host, port, user, passwd, db, conn, gormee}, nil
	}
}

func (ms Mysql) Host() string {
	return ms.host
}
func (ms Mysql) Port() int {
	return ms.port
}
func (ms Mysql) User() string {
	return ms.user
}
func (ms Mysql) Database() string {
	return ms.db
}
func (ms Mysql) Connection() *sql.DB {
	return ms.conn
}
func (ms *Mysql) Chost(host string) *Mysql {
	ms.host = host
	return ms
}
func (ms *Mysql) Chport(port int) *Mysql {
	ms.port = port
	return ms
}
func (ms *Mysql) Chusr(user string) *Mysql {
	ms.user = user
	return ms
}
func (ms *Mysql) Chpsd(old string, new string) *Mysql {
	if old == ms.passwd {
		ms.passwd = new
	}
	return ms
}
func (ms *Mysql) Chdb(db string) *Mysql {
	ms.db = db
	return ms
}

func (ms *Mysql) Connect(replace bool) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", ms.user, ms.passwd, ms.host, ms.port, ms.db))
	if replace {
		ms.Disconnect()
		if gormee, err := gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{}); err != nil {
			conn.Close()
			return nil, err
		} else {
			ms.conn = conn
			ms.gormee = gormee
			return conn, err
		}
	} else {
		return conn, err
	}
}

func (ms *Mysql) Disconnect() error {
	if ms.conn != nil {
		err := ms.conn.Close()
		ms.conn = nil
		ms.gormee = nil
		return err
	} else {
		return fmt.Errorf("connection is already closed")
	}
}

func (ms *Mysql) Query(query statement.Selector, constructor func() any) ([]any, error) {
	receiver := []any{}

	if rows, err := ms.conn.Query(query.String()); err != nil {
		return receiver, err
	} else {
		for rows.Next() {
			obj := constructor()
			ms.gormee.ScanRows(rows, &obj)
			receiver = append(receiver, obj)
		}
		rows.Close()
		return receiver, err
	}
}
func (ms *Mysql) Exec(query statement.Selector, args ...any) (sql.Result, error) {
	return ms.conn.Exec(query.String(), args...)
}
