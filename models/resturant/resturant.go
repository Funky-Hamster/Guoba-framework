package resturantmodel

type Resturant struct {
	Id          int32  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Discounts   string `db:"discounts" json:"discounts"`
	ImgUrl      string `db:"img_url" json:"img_url"`
	NumInAction int32  `db:"num_in_action" json:"num_in_action"`
}
