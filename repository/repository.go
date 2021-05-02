package repository

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Repository struct {
	engine *xorm.Engine
}

func New() *Repository {
	engine, err := xorm.NewEngine("postgres", "postgres://postgres:mypass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(2)

	return &Repository{
		engine: engine,
	}
}
