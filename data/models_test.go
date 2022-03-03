package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	db2 "github.com/upper/db/v4"
)

func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()
	defer fakeDB.Close()

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)
	models := fmt.Sprintf("%T", m)
	if models != "data.Models" {
		t.Error("Wrong type", models)
	}

	_ = os.Setenv("DATABASE_TYPE", "mysql")
	m = New(fakeDB)
	models = fmt.Sprintf("%T", m)
	if models != "data.Models" {
		t.Error("Wrong type", models)
	}

}

func Test_GetInsertID(t *testing.T) {
	var id db2.ID
	id = int64(1)

	returnedId := getInsertId(id)
	if fmt.Sprintf("%T", returnedId) != "int" {
		t.Error("wrong type returned")
	}

	id = 1
	returnedId = getInsertId(id)
	if fmt.Sprintf("%T", returnedId) != "int" {
		t.Error("wrong type returned")
	}

}
