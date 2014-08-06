package martinier

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
)

func NewServer(db *DatabaseConnection) *martini.Martini {
	engine := martini.New()
	engine.Use(render.Renderer(render.Options{IndentJSON: true}))
	engine.Use(db.Database())

	router := martini.NewRouter()

	router.Get("/users", binding.Json(User{}), func(user User, render render.Render, db *mgo.Database) {
		render.JSON(http.StatusOK, user.all(db))
	})
	router.Post("/users", binding.Json(User{}), func(user User, render render.Render, db *mgo.Database) {
		render.JSON(http.StatusCreated, user.store(db, user))
	})
	//router.Get("/users/:id", handler)
	//router.Put("/users/:id", handler)
	//router.Delete("/users/:id", handler)

	engine.Action(router.Handle)

	return engine
}
