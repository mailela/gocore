package db

import (
	"fmt"
	"log"

	"github.com/mailela/gorm"
	_ "github.com/mailela/gorm/dialects/sqlite"
	  _ "github.com/mailela/gorm/dialects/mysql"
         _ "github.com/mailela/gorm/dialects/postgres"
	 _ "github.com/mailela/gorm/dialects/sqlite"
	  _ "github.com/mailela/gorm/dialects/mssql"
	"github.com/mailela/gocore/EasyGoLib/utils"
)

type Model struct {
	ID        string         `structs:"id" gorm:"primary_key" form:"id" json:"id"`
	CreatedAt utils.DateTime `structs:"-" json:"createdAt" gorm:"type:datetime"`
	UpdatedAt utils.DateTime `structs:"-" json:"updatedAt" gorm:"type:datetime"`
	// DeletedAt *time.Time `sql:"index" structs:"-"`
}

var SQLite *gorm.DB

func Init() (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTablename string) string {
		return "t_" + defaultTablename
	}
	dbFile := utils.DBFile()
	log.Println("db file -->", utils.DBFile())
	SQLite, err = gorm.Open("sqlite", fmt.Sprintf("%s?loc=Asia/Shanghai", dbFile))
	if err != nil {
		return
	}
	// Sqlite cannot handle concurrent writes, so we limit sqlite to one connection.
	// see https://github.com/mattn/go-sqlite3/issues/274
	SQLite.DB().SetMaxOpenConns(1)
	SQLite.SetLogger(DefaultGormLogger)
	SQLite.LogMode(false)
	return
}

func Close() {
	if SQLite != nil {
		SQLite.Close()
		SQLite = nil
	}
}
