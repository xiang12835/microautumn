package models

//doc --- http://beego.me/docs/mvc/model/orm.md

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"microautumn/models/item"
	"microautumn/models/user"
)

/*
   run

   go run orm_sync.go orm syncdb
*/

func init() {

	fmt.Println("[init database]......")

	orm.Debug = true
	//regiter driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// register model
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(user.Post))
	orm.RegisterModel(new(item.Item))

	mysql_config := "root:@tcp(localhost:3306)/go_platform?charset=utf8"

	// set default database
	orm.RegisterDataBase("default", "mysql", mysql_config)
	//set db params

	orm.SetMaxIdleConns("default", 240)
	orm.SetMaxOpenConns("default", 240)

	// set go
	fmt.Println("[end init database]......")

}
