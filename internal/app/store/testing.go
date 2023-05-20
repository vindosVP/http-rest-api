package store

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"strings"
	"testing"
)

func TestStore(t *testing.T) (*Store, func(...string)) {
	t.Helper()
	cfg, err := config.NewConfig("../../../configs/apiserver_test.yaml")

	if err != nil {
		t.Fatal(err)
	}

	s := New(cfg, logrus.New())
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}

		}
		if err := s.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
