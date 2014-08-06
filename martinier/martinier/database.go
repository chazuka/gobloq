package martinier

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

type DatabaseConnection struct {
	*mgo.Session
	name string
}

func NewConnection(name string, dbServers string) *DatabaseConnection {
	connection, err := mgo.Dial(dbServers)
	if err != nil {
		panic(err)
	}
	ensureUserIndex(connection.DB(name))
	return &DatabaseConnection{connection, name}
}

func (d *DatabaseConnection) Database() martini.Handler {
	return func(context martini.Context) {
		s := d.Clone()
		context.Map(s.DB(d.name))
		defer s.Close()
		context.Next()
	}
}
