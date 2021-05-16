package application

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

type (
	ErrDuplicateKey struct {
		source             error
		ColumnOrConstraint string
	}
	ErrForeignKey struct {
		source             error
		ColumnOrConstraint string
	}
)

func firstNotEmpty(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

// tries to cast an error object as a PgError
// example.
// Return values:
// { *pgconn.PgError } Pointer to the error object
// { bool } If true, the object is a PgError
// Example:
/**
if pgErr, exists := AsPgError(err); exists {
	// do stuff
}
*/
func AsPgError(e error) (*pgconn.PgError, bool) {
	var pgErr *pgconn.PgError
	if errors.As(e, &pgErr) {
		return pgErr, true
	} else {
		return nil, false
	}
}

func ToAppError(pgError *pgconn.PgError) error {
	switch pgError.Code {
	case PgErrorUniqueViolation:
		return &ErrDuplicateKey{
			source:             pgError,
			ColumnOrConstraint: firstNotEmpty(pgError.ConstraintName, pgError.ColumnName),
		}
	case PgErrorForeignKeyViolation:
		return &ErrForeignKey{
			source:             pgError,
			ColumnOrConstraint: firstNotEmpty(pgError.ColumnName, pgError.ConstraintName),
		}
	default:
		return pgError
	}
}

func (e *ErrDuplicateKey) Error() string {
	return fmt.Sprintf("Duplicate Key Error: %v", e.source)
}

func (e *ErrDuplicateKey) Wrap() error {
	return e.source
}

func (e *ErrForeignKey) Error() string {
	return fmt.Sprintf("Foreign Key Error: %v", e.source)
}

func (e *ErrForeignKey) Wrap() error {
	return e.source
}
