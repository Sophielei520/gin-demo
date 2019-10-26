package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func main() {
	router := gin.Default()
	fmt.Println("router=", router)

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "graphql", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	fmt.Println("fields=", fields)
	fmt.Println("rootQuery=", rootQuery)

	//
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// fmt.Println("schemaConfig = ", schemaConfig)
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema ,error:%v", err)
	}
	query := `
	{
		hell
	}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql")
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s\n", rJSON)

}
