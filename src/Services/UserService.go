package Services

type UserService interface {
	all() [] Entities.User
	insert(Entities.User) Entities.User
}