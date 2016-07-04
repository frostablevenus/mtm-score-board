package main

import (
	//"encoding/json"
	/*"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"sort"*/

	"mtmScoreBoard/core/config"
	"mtmScoreBoard/core/routes"
	"mtmScoreBoard/resources"
)

func main() {
	cf := resources.ResourceConfig{
		IsEnablePostgres: true,
	}

	r, err := resources.Init(cf)
	if err != nil {
		return
	}
	//defer r.Close()

	engine := routes.CreateEngine(r)
	engine.Run(config.AppHost + ":" + config.AppPort)

	//
	//

	engine.Run()
}

/*

{"name" : "cam"}
{"name" : "John", "password" : "Cena"}
{"name" : "Spence", "password":"dude", "score":200}

*/
