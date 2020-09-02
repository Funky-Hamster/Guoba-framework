package conn

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
@func 创建db连接池
@author:柠檬191104
@date:2020/04/26
*/

var dbx *sqlx.DB

func init(){
	var err error
	dsn := "root:Guoba123456@tcp(cdb-fo0ulqw5.gz.tencentcdb.com:10137)/guoba"
	dbx,err = sqlx.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("init db conntection instalce failed:%v\n",err)
		return
	}
}

func NewDb() *sqlx.DB{
	if dbx != nil {
		return dbx
	}
	return  dbx
}

func CloseStmt(stmt *sql.Stmt){
	if stmt != nil {
		fmt.Println("close stmt connection")
		stmt.Close()
	}
}
