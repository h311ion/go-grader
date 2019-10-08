package user

type User struct {
	email string
	roles role
}

func NewUser(email string) User {
	return User{email: email, roles: role{"root", "admin", "gopher"}}
}

type role []string

type Profile struct {
	Email string
	Roles role
}

func GetProfile(user User) Profile {
	return Profile{
		user.email,
		role{
			"admin",
			"root",
		},
	}
}

func RegisterUser(email string, password string) error {
	// TODO Add user registration
	return nil
}
