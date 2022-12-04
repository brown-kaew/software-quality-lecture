package di

import (
	"database/sql"
	"testing"
)

type MockDB struct {
	query        string
	lastInsertId int64
	rowsAffected int64
}

func (md *MockDB) LastInsertId() (int64, error) {
	return md.lastInsertId, nil
}

func (md *MockDB) RowsAffected() (int64, error) {
	return md.rowsAffected, nil
}

func (md *MockDB) Exec(query string, args ...any) (sql.Result, error) {
	md.query = query
	return md, nil
}

func TestExecQuery(t *testing.T) {
	expectQuery := "SELECT * FROM sql"
	var expectRow int64 = 32

	mock := &MockDB{
		rowsAffected: expectRow,
	}

	row, _ := ExecQuery(mock, expectQuery)

	if expectQuery != mock.query {
		t.Errorf("Should have been call db.Exec wiith query but it not.")
	}

	if expectRow != row {
		t.Errorf("Should return row effect %d but it got %d.", expectRow, row)
	}
}
