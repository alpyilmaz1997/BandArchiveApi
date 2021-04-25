package repository

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Repository struct {
	engine *xorm.Engine
}

func New() *Repository {
	engine, err := xorm.NewEngine("postgres", "postgres://myapp:dbpass@localhost:1234/myapp")
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(2)

	return &Repository{
		engine: engine,
	}
}
