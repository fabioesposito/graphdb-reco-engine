package main

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
)

func main() {

	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()

	graph := rg.GraphNew("social", conn)
	query := `MATCH (p:person)-[v:visited]->(c:country)
           RETURN p.name, p.age, c.name`

	// result is a QueryResult struct containing the query's generated records and statistics.
	result, _ := graph.Query(query)

	// Pretty-print the full result set as a table.
	result.PrettyPrint()

	// Iterate over each individual Record in the result.
	fmt.Println("Visited countries by person:")
	for result.Next() { // Next returns true until the iterator is depleted.
		// Get the current Record.
		r := result.Record()

		// Entries in the Record can be accessed by index or key.
		pName := r.GetByIndex(0)
		fmt.Printf("\nName: %s\n", pName)
		pAge, _ := r.Get("p.age")
		fmt.Printf("\nAge: %d\n", pAge)
	}

	// Path matching example.
	query = "MATCH p = (:person)-[:visited]->(:country) RETURN p"
	result, err := graph.Query(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Pathes of persons visiting countries:")
	for result.Next() {
		r := result.Record()
		p, ok := r.GetByIndex(0).(rg.Path)
		fmt.Printf("%s %v\n", p, ok)
	}

}
