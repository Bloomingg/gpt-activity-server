package main

import (
	"customer/conf"
	"customer/db/mysql"
	"customer/route"

	logging "github.com/sirupsen/logrus"
)

func main() {
	r := route.NewRoute()
	err := r.Run(conf.HttpPort)
	if err != nil {
		logging.Fatalln(err)
	}
}

func init() {
	conf.Init("./conf/config.ini")
	mysql.Init()
}
