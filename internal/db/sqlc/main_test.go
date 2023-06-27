package sqlc

import (
	"database/sql"
	"github.com/borntodie-new/backend-master-class/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testStore Store
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	log.Printf("%v\n", config)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testStore = NewStore(testDB)
	os.Exit(m.Run())
}
