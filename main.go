package main

import (
	"fmt"

	"github.com/kagi-opendoor222/go_practice3/app/controllers"
	"github.com/kagi-opendoor222/go_practice3/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
