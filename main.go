package main

import (
	"context"

	"go.uber.org/fx"

	"github.com/nahueldev23/inventory-go/database"
	"github.com/nahueldev23/inventory-go/internal/repository"
	"github.com/nahueldev23/inventory-go/internal/service"
	"github.com/nahueldev23/inventory-go/settings"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "nahuel@gmail.con", "nahuel", "123456")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "nahuel@gmail.com", "123456")
				if err != nil {
					panic(err)
				}

				if u.Name != "nahuel" {
          
					panic(err)
				}

			},
		),
	)

	app.Run()
}
