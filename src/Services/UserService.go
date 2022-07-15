package Services

import "github.com/abdalrazzak/gin-golang-test/Entities"

type UserService interface {
	all() [] Entities.User
	insert(Entities.User) Entities.User
}