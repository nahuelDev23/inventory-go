package main

import (
  "context"
  "go.uber.org/fx"

  "github.com/nahueldev23/inventory-go/database"
  "github.com/nahueldev23/inventory-go/settings"
  "github.com/jmoiron/sqlx"
)

func main()  {
  app := fx.New(
    fx.Provide(
      context.Background,
      settings.New,
      database.New,
      ),
    fx.Invoke(
      func(db *sqlx.DB){
        _,err := db.Query("SELECT * FROM USERS")
        if err != nil {
          panic(err)
        }
      },
      ),
  )

  app.Run()
}
