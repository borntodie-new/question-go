package main

import (
	"database/sql"
	"github.com/borntodie-new/question-go/internal/config"
	"github.com/borntodie-new/question-go/internal/handlers"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cnf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}
	conn, err := sql.Open(cnf.DBDriver, cnf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	core, err := handlers.NewCore(conn, cnf)
	if err != nil {
		log.Fatalln("create core object fail: ", err)
	}
	err = core.Run()
	if err != nil {
		log.Fatalln("start application failed: ", err)
	}

}
