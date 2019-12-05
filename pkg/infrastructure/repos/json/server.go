package json

import "errors"

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionScout identifier for the JSON collection of beers
	CollectionScout = "scouts"
)

//Repo represents the storage of Scout entities
// Storage stores beer data in JSON files
type Repo struct {
	db *scribble.Driver
}

// Create etc
func (r *Repo) Create(scout Scout) (string, error) {
	return "mockup", nil
}

// Read etc
func (r *Repo) Read(scoutID string) (Scout, error) {
	scout := new(Scout)
	scout.ID = "mpoling"
	scout.Name = "Mark Poling"
	return scout, nil
}

// Update etc
func (r *Repo) Update(scout Scout) error {
	return errors.New("Test of Update error catching")
}

// Delete etc
func (r *Repo) Delete(scoutID string) error {
	return errors.New("Test of Delete error catching")
}
