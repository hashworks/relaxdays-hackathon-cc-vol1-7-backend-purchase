package main

import (
	_ "embed"
	"flag"
	"log"

	"database/sql"

	sqlite "github.com/mattn/go-sqlite3"

	"github.com/agnivade/levenshtein"
	"github.com/gchaincl/dotsql"
	_ "github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-purchase/docs"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-purchase/router"
)

//go:embed sql/init.sql
var sqlInit string

//go:embed sql/alter.sql
var sqlAlter string

//go:embed sql/select.sql
var sqlSelect string

// @title Purchase Backend Task
// @version 1.0
// @description Solution for 'Einkauf' backend task of https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite#h.klg8hathdmsn

// @contact.name Justin Kromlinger
// @contact.url https://hashworks.net
// @contact.email justin.kromlinger@stud.htwk-leipzig.de

// @license.name GNU Affero General Public License v3
// @license.url https://gnu.org/licenses/agpl.html

// @host 127.0.0.1:8080
// @BasePath /
func main() {
	dsn := flag.String("dsn", "file::memory:?cache=shared", "SQLite database DSN")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("Missing database dsn")
	}

	var server router.Server
	var err error

	// Load init commands
	dotInit, err := dotsql.LoadFromString(sqlInit)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Load alter commands
	server.DotAlter, err = dotsql.LoadFromString(sqlAlter)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Load select commands
	server.DotSelect, err = dotsql.LoadFromString(sqlSelect)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Register custom sql functions
	sql.Register("sqlite3_custom", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("levenshteinDistance", levenshtein.ComputeDistance, true); err != nil {
				return err
			}
			return nil
		}})

	// Open database
	server.DB, err = sql.Open("sqlite3_custom", *dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer server.DB.Close()

	// Init database
	for _, command := range []string{
		"create-table-purchase",
	} {
		_, err := dotInit.Exec(server.DB, command)
		if err != nil {
			log.Fatalf("Failed to init database: %s failed with %s\n", command, err.Error())
		}
	}

	routerEngine := server.NewRouter()

	log.Println("Starting purchase backend on 0.0.0.0:8080")

	log.Fatal(routerEngine.Run(":8080"))
}
