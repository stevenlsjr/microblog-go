package application

import (
	"errors"
	"github.com/go-playground/assert"
	"github.com/jackc/pgconn"
	"testing"
)

func TestFirstNotEmpty(t *testing.T) {
	assert.Equal(t, firstNotEmpty("", "bar", "foo"), "bar")
	assert.Equal(t, firstNotEmpty("foo", "", "bar"), "foo")
	assert.Equal(t, firstNotEmpty("", "", ""), "")
	assert.Equal(t, firstNotEmpty("", "", " "), " ")
}

func AssertErrorIs(t *testing.T, e error, expected error) {
	if !errors.As(e, &expected) {
		t.Errorf("error %v is not %v", e, expected)
	}
}

func TestToAppError(t *testing.T) {
	type testCase struct {
		pgError    *pgconn.PgError
		expectType error
	}
	testCases := []testCase{
		{
			pgError:    &pgconn.PgError{Code: PgErrorForeignKeyViolation, ColumnName: "foo"},
			expectType: &ErrForeignKey{},
		},
		{
			pgError:    &pgconn.PgError{Code: PgErrorUniqueViolation, ColumnName: "foo"},
			expectType: &ErrDuplicateKey{},
		},
		{ // unexpected error
			pgError:    &pgconn.PgError{Code: PgErrorAdminShutdown},
			expectType: &pgconn.PgError{Code: PgErrorAdminShutdown},
		},
	}

	for _, testCase := range testCases {
		err := ToAppError(testCase.pgError)
		AssertErrorIs(t, err, testCase.expectType)
	}
}
