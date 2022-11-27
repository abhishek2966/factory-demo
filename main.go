package main

import (
	"fmt"

	"github.com/abhishek2966/factory-demo/pkg/pgdb"
)

func main() {
	basedb1 := "pgdb.db1"
	db1Source := pgdb.NewDBSource(basedb1)

	//new db1 connection is opened
	db1, err := pgdb.OpenConnection(db1Source)
	// do some query work on db1 connection
	fmt.Println(db1, err)

	basedb2 := "pgdb.db2"
	db2Source := pgdb.NewDBSource(basedb2)

	//other db2 connection is opened
	db2, err := pgdb.OpenConnection(db2Source)
	// do some query work on db2 connection
	fmt.Println(db2, err)
}
