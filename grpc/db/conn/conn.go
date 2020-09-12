package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

func init() {
	var err error
	dsn := "root:Guoba123456@tcp(cdb-fo0ulqw5.gz.tencentcdb.com:10137)/guoba"
	dbx, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("init db conntection instance failed:%v\n", err)
		return
	}
}

func NewDb() *sqlx.DB {
	if dbx != nil {
		return dbx
	}
	return dbx
}

func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		fmt.Println("close stmt connection")
		stmt.Close()
	}
}
