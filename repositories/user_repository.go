package repositories

type UserRepository interface {
	ValidateUser(username, password string) (bool, error)
	CreateUser(username, password string) (bool, error)
	DeleteUser(username string) (bool, error)
}
