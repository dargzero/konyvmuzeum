package main

import (
	"log"
	"net/http"

	"okki.hu/konyvmuzeum/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":13400", api.NewRouter()))
}

// client, err := mongo.NewClient(
// 	options.Client().ApplyURI("mongodb://db:27017"))
// if err != nil {
// 	log.Fatal(err)
// }
// ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
// err = client.Connect(ctx)
// if err != nil {
// 	log.Fatal(err)
// }
// defer client.Disconnect(ctx)
// names, err := client.ListDatabaseNames(ctx, bson.D{})
// if err != nil {
// 	log.Fatal(err)
// }
// for _, name := range names {
// 	fmt.Println(name)
// }
