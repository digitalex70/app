package data

import (
	"errors"
	"time"

	up "github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Active    int       `db:"user_active"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetAll() ([]*User, error) {
	collection := upper.Collection(u.Table())
	var all []*User
	res := collection.Find().OrderBy("last_name")
	err := res.All(&all)
	if err != nil {
		return nil, err
	}
	return all, err
}

func (u *User) GetByEmail(email string) (*User, error) {
	collection := upper.Collection(u.Table())
	var theUser User
	res := collection.Find(up.Cond{"email =": email})
	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	var token Token
	collection = upper.Collection(token.Table())
	res = collection.Find(up.Cond{"user_id =": theUser.Id, "expiry >": time.Now()}).OrderBy("created_at desc")
	err = res.One(&token)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}
	theUser.Token = token
	return &theUser, err
}

func (u *User) Get(id int) (*User, error) {
	collection := upper.Collection(u.Table())
	var theUser User
	res := collection.Find(up.Cond{"id =": id})
	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	//var token Token
	//collection = upper.Collection(token.Table())
	//res = collection.Find(up.Cond{"user_id =": theUser.Id, "expiry >": time.Now()}).OrderBy("created_at desc")
	//err = res.One(&token)
	//if err != nil {
	//	if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
	//		return nil, err
	//	}
	//}
	//theUser.Token = token
	return &theUser, err
}

func (u *User) Update(user User) error {
	user.UpdatedAt = time.Now()
	collection := upper.Collection(u.Table())
	res := collection.Find(user.Id)
	err := res.Update(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(id int) error {
	collection := upper.Collection(u.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Insert(user User) (int, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	user.Password = string(newHash)

	collection := upper.Collection(u.Table())
	res, err := collection.Insert(user)
	if err != nil {
		return 0, err
	}
	id := getInsertId(res.ID())
	return id, nil
}

func (u *User) ResetPassword(id int, password string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user, err := u.Get(id)
	if err != nil {
		return err
	}
	user.Password = string(newHash)
	err = user.Update(*user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
