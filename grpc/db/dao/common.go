package dao

import (
	"github.com/gin-gonic/gin/examples/grpc/db/conn"
	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

func init() {
	dbConn = conn.NewDb()
}
