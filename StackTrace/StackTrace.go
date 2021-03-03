package StackTrace

import (
	"github.com/enigs/go-utilities/DatabaseUtils"
	"github.com/guregu/null"
	"gorm.io/gorm"
)

// Set StackTrace struct
type StackTrace struct {
	ID                  null.String		`gorm:"primaryKey"`
	FunctionName		null.String
	FileName 			null.String
	FileLine            null.String
	Error               null.String
	CreatedAt           null.Time
	UpdatedAt 			null.Time
}

// Creates a constructor for StackTrace
func NewStackTrace() *StackTrace {
	// Set StackTrace struct
	var stacktrace StackTrace

	// Return StackTrace
	return &stacktrace
}

// Check if settings has table
func (stacktrace *StackTrace) HasTable(store *gorm.DB) bool {
	return store.Migrator().HasTable(stacktrace)
}

// Create table
func (stacktrace *StackTrace) CreateTable(store *gorm.DB) error {
	// Create query
	query := `CREATE TABLE stack_traces (
	  id varchar(36) NOT NULL,
	  function_name varchar(255) DEFAULT NULL,
	  file_name varchar(255) DEFAULT NULL,
	  file_line varchar(255) DEFAULT NULL,
	  error varchar(255) DEFAULT NULL,
	  created_at timestamp WITH TIME ZONE DEFAULT NULL,
	  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	ALTER TABLE stack_traces ADD PRIMARY KEY (id);`

	// Execute query
	result := store.Exec(query)

	// Check if query has error
	if DatabaseUtils.HasError(result.Error) {
		return result.Error
	}

	return nil
}

