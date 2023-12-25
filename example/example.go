package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/jodydadescott/openhab-go-sdk"
)

func main() {

	client := sdk.New(&sdk.Config{
		// checkov:skip=CKV_SECRET_6: ADD REASON
		APIToken: os.Getenv("OPENHAB_TOKEN"),
		API:      os.Getenv("OPENHAB_API"),
	})

	things, err := client.GetThings(context.Background())
	if err != nil {
		panic(err)
	}

	for _, thing := range *things {
		fmt.Println(thing.Properties.ServiceName)
	}
}
