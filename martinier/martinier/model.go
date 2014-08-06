package martinier

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Email     string        `json:"email"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func ensureUserIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
	}
	indexErr := db.C("users").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}

func (model *User) all(db *mgo.Database) []User {
	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}

func (model *User) single(db *mgo.Database, id bson.ObjectId) *User {
	user := User{}
	err := db.C("users").Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		panic(err)
	}

	return &user
}

func (model *User) store(db *mgo.Database, binded User) *User {
	binded.ID = bson.NewObjectId()
	binded.CreatedAt = time.Now().Local()
	binded.UpdatedAt = time.Now().Local()
	err := db.C("users").Insert(binded)
	if err != nil {
		panic(err)
	}

	return model.single(db, binded.ID)
}
