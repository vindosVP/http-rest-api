package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"strings"
	"testing"
)

func TestDB(t *testing.T) (*sql.DB, func(...string)) {
	t.Helper()
	conf, err := config.NewConfig("../../../../configs/apiserver_test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	err = logger.ConfigureLogger(conf.LogLevel)
	if err != nil {
		t.Fatal(err)
	}

	cStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.DBName)

	db, err := sql.Open("postgres", cStr)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		err := db.Close()

		if err != nil {
			t.Fatal(err)
		}

	}

}
