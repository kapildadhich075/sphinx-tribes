package db

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var now = time.Now()

var code = ConnectionCodes{
	ID:               1,
	ConnectionString: "2222222",
	IsUsed:           false,
	DateCreated:      &now,
}

func TestCodeInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	prep := mock.ExpectPrepare("^INSERT INTO connectioncodes*")

	prep.ExpectExec().
		WithArgs(code.ID, code.ConnectionString, code.IsUsed, code.DateCreated).
		WillReturnResult(sqlmock.NewResult(int64(code.ID), 1))

	mock.ExpectCommit()

	assert.Nil(t, err)
}

func TestCodeGet(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"connection_string", "date_created", "is_used", "date_created"}).AddRow(code.ID, code.ConnectionString, code.IsUsed, code.DateCreated)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT connection_string, date_created FROM connectioncodes WHERE is_used = ? ORDER BY id DESC LIMIT 1`)).
		WithArgs(false).
		WillReturnRows(rows)

	assert.Nil(t, err)
}
