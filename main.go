package main

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
)

type RootResolver struct{}

type Account struct {
	ID        graphql.ID
	FirstName string
	LastName  string
	Email     string
}

var (
	opts     = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	accounts = []Account{
		{
			ID:        "12345",
			FirstName: "fuga",
			LastName:  "hoge",
			Email:     "sample@example.com",
		},
	}
)

func (r *RootResolver) Accounts() ([]Account, error) {
	return accounts, nil
}

func (r *RootResolver) CreateAccount(args struct {
	FirstName string
	LastName  string
	Email     string
}) (Account, error) {
	newAccount := Account{
		ID:        graphql.ID(fmt.Sprint(len(accounts)) + "-account"),
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Email:     args.Email,
	}

	accounts = append(accounts, newAccount)
	return newAccount, nil
}

// Reads and parses the schema from file.
// Associates root resolver. Panics if can't read.
func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return parsedSchema
}

func main() {
	http.Handle("/graphql", &relay.Handler{
		Schema: parseSchema("./schema.graphql", &RootResolver{}),
	})

	fmt.Println("serving on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
