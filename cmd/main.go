package main

import (

	//"io/ioutil"
	// "encoding/json"
	// "github.com/mbatimel/WB_Tech-L0/tree/main/iternal/Data"
	"github.com/mbatimel/WB_Tech-L0/tree/main/iternal/Server"
	// "github.com/mbatimel/WB_Tech-L0/tree/main/iternal/migrate"
)

func main() {
    srv := server.Server{}
    if err := srv.Up(); err != nil {
        panic(err)
    }
}
