package main

import (
  "go.uber.org/fx"
  "github.com/nahueldev23/inventory-go/settings"
  "log"
)

func main()  {
  app := fx.New(
    fx.Provide(
      settings.New,
      ),
    fx.Invoke(),
  )

  app.Run()
}
