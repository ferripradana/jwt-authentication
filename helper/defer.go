package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		IfErrorPanic(errRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		IfErrorPanic(errorCommit)
	}

}
