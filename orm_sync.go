package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "microautumn/models"
)

func main() {
	fmt.Println("starting......")
	orm.RunCommand()
	//orm.RunSyncdb("default", false, true)
	fmt.Println("ended......")
}
