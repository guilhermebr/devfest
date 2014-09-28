package user

import "fmt"

type User struct {
	ID        int64
	Nome      string
	Sobrenome string
	Idade     int
}

func NewUser(nome string, sobrenome string, idade int) (*User, error) {
	if nome == "" {
		return nil, fmt.Errorf("nome obrigatorio")
	}
	return &User{0, nome, sobrenome, idade}, nil
}

type UserManager struct {
	users  []*User
	lastID int64
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (m *UserManager) Save(user *User) error {
	if user.ID == 0 {
		m.lastID++
		user.ID = m.lastID
		m.users = append(m.users, cloneUser(user))
		return nil
	}

	for i, u := range m.users {
		if u.ID == user.ID {
			m.users[i] = cloneUser(user)
			return nil
		}
	}
	return fmt.Errorf("unknown user")
}

func (m *UserManager) All() []*User {
	return m.users
}

func (m *UserManager) Find(ID int64) (*User, bool) {
	for _, u := range m.users {
		if u.ID == ID {
			return u, true
		}
	}
	return nil, false
}

//Help Functions
func cloneUser(u *User) *User {
	c := *u
	return &c
}
