package db

import (
	// "fmt"
	"log"
	"time"

	"github.com/upper/db/v4/adapter/sqlite"
)

var settings = sqlite.ConnectionURL{
	Database: `example.db`, // Path to database file
}

type Birthday struct {
	// The 'name' column of the 'birthday' table
	// is mapped to the 'name' property.
	Name string `db:"name"`
	// The 'born' column of the 'birthday' table
	// is mapped to the 'born' property.
	Born time.Time `db:"born"`
}

func All() []Birthday {
	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()
	birthdayCollection := sess.Collection("birthday")

	res := birthdayCollection.Find()

	var birthdays []Birthday

	err = res.All(&birthdays)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	// r2 := make(map[string]string)
	// for _, birthday := range birthdays {
	// 	r2[birthday.Name] = birthday.Born.Format("January 2, 2006")
	// }
	// return r2
	return birthdays
}

// func main() {

// 	// Attempt to open the 'example.db' database file
// 	sess, err := sqlite.Open(settings)
// 	if err != nil {
// 		log.Fatalf("db.Open(): %q\n", err)
// 	}
// 	defer sess.Close() // Closing the session is a good practice.

// 	// The 'birthday' table is referenced.
// 	birthdayCollection := sess.Collection("birthday")

// 	// Any rows that might have been added between the creation of
// 	// the table and the execution of this function are removed.
// 	err = birthdayCollection.Truncate()
// 	if err != nil {
// 		log.Fatalf("Truncate(): %q\n", err)
// 	}

// 	// Three rows are inserted into the 'birthday' table.
// 	birthdayCollection.Insert(Birthday{
// 		Name: "Hayao Miyazaki",
// 		Born: time.Date(1941, time.January, 5, 0, 0, 0, 0, time.Local),
// 	})

// 	birthdayCollection.Insert(Birthday{
// 		Name: "Nobuo Uematsu",
// 		Born: time.Date(1959, time.March, 21, 0, 0, 0, 0, time.Local),
// 	})

// 	birthdayCollection.Insert(Birthday{
// 		Name: "Hironobu Sakaguchi",
// 		Born: time.Date(1962, time.November, 25, 0, 0, 0, 0, time.Local),
// 	})

// 	// The database is queried for the rows inserted.
// 	res := birthdayCollection.Find()

// 	// The 'birthdays' variable is filled with the results found.
// 	var birthdays []Birthday

// 	err = res.All(&birthdays)
// 	if err != nil {
// 		log.Fatalf("res.All(): %q\n", err)
// 	}

// 	// The 'birthdays' variable is printed to stdout.
// 	for _, birthday := range birthdays {
// 		fmt.Printf("%s was born in %s.\n",
// 			birthday.Name,
// 			birthday.Born.Format("January 2, 2006"),
// 		)
// 	}
// }
