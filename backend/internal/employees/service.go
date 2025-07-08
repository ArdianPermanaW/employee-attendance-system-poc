package employees

type Service interface {
	Create(emp *Employee) error
	GetAll() ([]Employee, error)
	GetByID(id uint) (*Employee, error)
	Update(emp *Employee) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(emp *Employee) error         { return s.repo.Create(emp) }
func (s *service) GetAll() ([]Employee, error)        { return s.repo.FindAll() }
func (s *service) GetByID(id uint) (*Employee, error) { return s.repo.FindByID(id) }
func (s *service) Update(emp *Employee) error         { return s.repo.Update(emp) }
func (s *service) Delete(id uint) error               { return s.repo.Delete(id) }
