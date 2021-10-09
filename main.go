package main

import (
	"fmt"
	"log"

	"github.com/kagi-opendoor222/go_practice3/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.Logfile)

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "testman"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	u, err := models.GetUser(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)
}
