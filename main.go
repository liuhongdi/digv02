package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liuhongdi/digv02/global"
	"github.com/liuhongdi/digv02/router"
	"log"
)

//init
func init() {
	err := global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}




