package passwordless

import (
	"database/sql"
	log "github.com/Sirupsen/logrus"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	stdlog "log"
	"os"
)

var dbmap *gorp.DbMap

func init() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Panic("Missing required environment variable 'DATABASE_URL'.")
	}

	db, err := sql.Open("postgres", dbUrl)
	if nil != err {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Database connection error")
	}

	dbmap = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}

	if os.Getenv("DEBUG") == "true" {
		dbmap.TraceOn("[gorp]", stdlog.New(os.Stdout, "passwordless:", stdlog.Lmicroseconds))
	}

	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln("Create tables failed", err)
	}
	//checkErr(err, "Create tables failed")
}
