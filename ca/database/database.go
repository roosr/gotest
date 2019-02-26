package cadb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "digicert"
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}

var (
	ErrOk        = fmt.Errorf("ok")
	ErrDb        = fmt.Errorf("db_connection")
	ErrMigration = fmt.Errorf("no_migration")
)

func RunSchemaMigration() error {

	db := dbConn()
	if db == nil {
		return ErrDb
	}

	//filesData := migration.Files()

	dbDriver, _ := mysql.WithInstance(db, &mysql.Config{})

	var files file.File
	srcDriver, _ := files.Open("file://c:/data/migrations")

	migrater, migraterErr := migrate.NewWithInstance("source", srcDriver,
		"dbname", dbDriver)

	if migraterErr != nil {
		if migraterErr == migrate.ErrLocked {
			return nil
		}

		fmt.Print(migraterErr.Error())
		return ErrMigration
	}

	migraterErr = migrater.Up()
	if migraterErr != nil {
		if migraterErr == migrate.ErrNoChange {
			return nil
		}

		fmt.Print(migraterErr.Error())
		return ErrMigration
	}

	return nil
}
