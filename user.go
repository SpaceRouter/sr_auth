package sr_auth

type Role string

type User struct {
	Login     string
	FirstName string
	LastName  string
	Email     string
	Roles     []Role
}

func (u *User) HasRole(role Role) bool {
	for _, uRole := range u.Roles {
		if role == uRole {
			return true
		}
	}
	return false
}