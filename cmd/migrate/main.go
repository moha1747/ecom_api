package migrate

import (
	"log"
	"os"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/moha1747/ecom_api/config"
	"github.com/moha1747/ecom_api/db"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	// remove the extra closing brace
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migratemigrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args) - 1]
	if cmd == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
			if err := m.Down(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
