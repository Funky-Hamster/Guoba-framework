package resturantmodel

type IResturant interface {
	AddResturant(r *Resturant) (int64, error)

	// UpdateResturant(r *Resturant) (int64, error)

	// DeleteResturantById(id int) (int64, error)

	// GetResturantById(id int) (Resturant, error)

	// ListResturants(limit, page int) ([]Resturant, error)

	GetResturants() ([]*Resturant, error)
}
