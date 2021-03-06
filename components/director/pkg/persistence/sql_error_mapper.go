package persistence

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/kyma-incubator/compass/components/director/pkg/resource"

	"github.com/kyma-incubator/compass/components/director/pkg/apperrors"

	"github.com/lib/pq"
)

func MapSQLError(err error, resourceType resource.Type, sqlOperation resource.SQLOperation, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		log.Errorf("SQL: no rows in result set for '%s' resource type", resourceType)
		return apperrors.NewNotFoundErrorWithType(resourceType)
	}

	pgErr, ok := err.(*pq.Error)
	if !ok {
		log.Errorf("Error while casting to postgres error. Actual error: %s", err)
		return apperrors.NewInternalError("Unexpected error while executing SQL query")
	}

	log.Errorf("SQL Error: %s. Caused by: %s. DETAILS: %s", fmt.Sprintf(format, args...), pgErr.Message, pgErr.Detail)

	switch pgErr.Code {
	case NotNullViolation:
		return apperrors.NewNotNullViolationError(resourceType)
	case CheckViolation:
		return apperrors.NewCheckViolationError(resourceType)
	case UniqueViolation:
		return apperrors.NewNotUniqueError(resourceType)
	case ForeignKeyViolation:
		return apperrors.NewForeignKeyInvalidOperationError(sqlOperation, resourceType)
	}

	return apperrors.NewInternalError("Unexpected error while executing SQL query")
}
