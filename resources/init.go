package resources

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"mtmScoreBoard/resources/constants"
)

type ResourceConfig struct {
	IsEnablePostgres bool
}

type Resource struct {
	Config     ResourceConfig
	PostgreSql *sql.DB
}

func (r *Resource) Close() {
	if r.Config.IsEnablePostgres {
		r.PostgreSql.Close()
	}
}

func initPostgreSQL() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://"+constants.PostgresUser+
		":"+constants.PostgresPassword+"@"+constants.PostgresHost+
		":"+constants.PostgresPort+"/"+constants.PostgresDB+
		"?sslmode=disable")
	//postgres://postgres:postgres5263@127.0.0.1:5432/mtmdb?sslmode=disable

	return db, err
}

func Init(config ResourceConfig) (*Resource, error) {
	r := Resource{}
	r.Config = config

	if config.IsEnablePostgres {
		db, err := initPostgreSQL()
		if err != nil {
			fmt.Println("Connect PostgreSQL Failed...", err)
			return nil, err
		}
		r.PostgreSql = db

		fmt.Printf("%T %v\n", db, err)
		fmt.Println(db == nil)
		fmt.Println(r.PostgreSql == nil)
	}

	return &r, nil
}
