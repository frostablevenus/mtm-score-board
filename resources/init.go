package resources

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"mtm-score-board/resources/constants"
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
	}

	return &r, nil
}
