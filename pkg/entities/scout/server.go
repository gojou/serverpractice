package scout

//Scout defines the fields for the Scout server
type Scout struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Server interface
type Server interface {
	Create(scout Scout) (string, error)
	Read(scoutID string) (Scout, error)
	Update(scout Scout) error
	Delete(scoutID string) error
}

// Repo interface
type Repo interface {
	Create(scout Scout) (string, error)
	Read(scoutID string) (Scout, error)
	Update(scout Scout) error
	Delete(scoutID string) error
}

type server struct {
	r Repo
}

//NewServer returns the address to a Scout server
func NewServer(r Repo) Server {
	return &server{r}
}

func (s *server) Create(scout Scout) (string, error) {
	return s.r.Create(scout)
}

func (s *server) Read(scoutID string) (Scout, error) {
	return s.r.Read(scoutID)
}
func (s *server) Update(scout Scout) error {
	return s.r.Update(scout)
}
func (s *server) Delete(scoutID string) error {
	return s.r.Delete(scoutID)
}
