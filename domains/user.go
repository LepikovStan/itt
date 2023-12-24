package domains

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	firstname string
	lastname  string
	fullname  string
	age       int
	isMarried bool
	password  string
}

func (u User) SetFirstName(n string) User {
	u.firstname = n
	return u
}

func (u User) SetLastName(n string) User {
	u.lastname = n
	return u
}

func (u User) SetAge(age int) User {
	u.age = age
	return u
}

func (u User) SetIsMarried(b bool) User {
	u.isMarried = b
	return u
}

func (u User) SetPassword(p string) User {
	u.password = p
	return u
}

func (u User) SetID(id uuid.UUID) User {
	u.id = id
	return u
}

func NewUser() User {
	return User{
		id: uuid.New(),
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Firstname() string {
	return u.firstname
}

func (u User) Lastname() string {
	return u.firstname
}

func (u User) Fullname() string {
	return u.firstname + " " + u.lastname
}

func (u User) Age() int {
	return u.age
}

func (u User) IsMarried() bool {
	return u.isMarried
}
