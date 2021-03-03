package StackTrace

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/guregu/null"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

// Save stack trace in database
func (stacktrace *StackTrace) Create(store *gorm.DB, err error)  {
	sk := 1

	var pc uintptr
	var file string
	var line int
	var ok bool

	pc, file, line, ok = runtime.Caller(sk)

	if ok {
		fn := runtime.FuncForPC(pc)
		name := fn.Name()
		ix := strings.LastIndex(name, ".")

		if ix > 0 && (ix+1) < len(name) {
			name = name[ix+1:]
		}

		nd, nf := filepath.Split(file)

		// Check if Database has settings table
		hasStackTraceTable := stacktrace.HasTable(store)

		if !hasStackTraceTable {
			_ = stacktrace.CreateTable(store)
		}

		stacktrace.ID = null.NewString(ksuid.New().String(), true)
		stacktrace.FunctionName = null.NewString(fmt.Sprintf("Function Name: %s", name), true)
		stacktrace.FileName = null.NewString(fmt.Sprintf("File Name: %s",filepath.Join(filepath.Base(nd), nf)), true)
		stacktrace.FileLine = null.NewString(fmt.Sprintf("File Line: %d", line), true)
		stacktrace.CreatedAt = null.NewTime(time.Now().UTC(), true)
		stacktrace.UpdatedAt = null.NewTime(time.Now().UTC(), true)


		if err != nil {
			stacktrace.Error = null.NewString(err.Error(), true)
		}

		if err != nil {
			stacktrace.Error = null.NewString(err.Error(), true)
		}

		store.Create(stacktrace)
	}

}