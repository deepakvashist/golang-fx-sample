package internal

// Usecase represents the application user usecase interface.
type Usecase interface {
	GetByID(id int) (*User, error)
}

// NewUsecase is responsible by building and returninig an implementation of Usecase interface.
func NewUsecase(repository Repository) Usecase {
	return &usecase{repository: repository}
}

type usecase struct {
	repository Repository
}

// GetByID returns a user entity by ID.
func (u usecase) GetByID(id int) (*User, error) {
	return u.repository.GetByID(id)
}
