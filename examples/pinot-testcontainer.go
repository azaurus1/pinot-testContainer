package examples

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/azaurus1/go-pinot-api/model"
)

func CreateUser() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancel()

	container := pinot_testContainer.NewPinotContainer()
	pinot, err := container.RunPinotContainer(ctx)

	user := model.User{
		Username:  "testUser",
		Password:  "password",
		Component: "BROKER",
		Role:      "admin",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	createResp, err := pinot.CreateUser(ctx, userBytes)
	pinot.TearDown()
}
