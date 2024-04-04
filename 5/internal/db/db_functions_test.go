package db

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

type MockDB struct{}

func (m MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if query == "SELECT name FROM users" {
		return nil, errors.New("mock error")
	}
	return nil, nil
}

func TestDBService_GetNames(t *testing.T) {
	mockDB := MockDB{}
	dbService := New(mockDB)

	names, err := dbService.GetNames()

	if !reflect.DeepEqual(names, []string{}) {
		t.Errorf("Expected empty array, but got %v", names)
	}

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
