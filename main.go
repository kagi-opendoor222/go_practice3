package main

import (
	"fmt"

	"github.com/kagi-opendoor222/go_practice3/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.Logfile)

	fmt.Println(models.Db)
}
