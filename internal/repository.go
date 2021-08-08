package internal

// Repository represents the application user repository interface.
type Repository interface {
	GetByID(id int) (*User, error)
}

// NewRepository is responsible by building and returninig an implementation of Repository interface.
func NewRepository() Repository {
	return &repository{}
}

type repository struct{}

// GetByID returns a mocked user entity by ID.
func (r repository) GetByID(id int) (*User, error) {
	return &User{ID: 1, Email: "test@test.com", Username: "test"}, nil
}
