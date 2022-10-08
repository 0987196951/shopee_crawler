package crawler

type Category struct {
	Catid            int64      `json:"catid"`
	Parent_catid     int64      `json:"parent_catid"`
	Name             string     `json:"name"`
	Display_name     string     `json:"display_name"`
	Image            string     `json:"image"`
	Unselected_image string     `json:"unselected_image"`
	Selected_image   string     `json:"selected_image"`
	Level            int        `json:"level"`
	Children         []Category `json:"children"`
}

type Product struct {
	Name            string          `json:"name" bson:"name,omitempty"`
	Description     string          `json:"description" bson:"description,omitempty"`
	ProductID       string          `json:"productID" bson:"product_id,omitempty"`
	Image           string          `json:"image" bson:"image,omitempty"`
	UrlProduct      string          `json:"url" bson:"product_url,omitempty"`
	Price           string          `json:"price" bson:"price,omitempty"`
	AggregateRating AggregateRating `json:"aggregateRating" bson:"aggregate_rating,omitempty"`
	Offers          Offers          `json:"offers" bson:"offers,omitempty"`
	Sold            int             `bson:"sold,omitempty"`
	Stock           int             `bson:"stock,omitempty"`
	ShipsFrom       string          `bson:"ship_from,omitempty"`
	Category        string          `bson:"category,omitempty"`
	Options         []string        `bson:"options,omitempty"`
}
type Offers struct {
	Price     string `json:"price" bson:"price,omitempty"`
	Seller_id int64  `bson:"seller_id,omitempty"`
}
type AggregateRating struct {
	BestRating  int    `json:"bestRating" bson:"best_rating,omitempty"`
	WorstRating int    `json:"worstRating" bson:"worst_rating,omitempty"`
	RatingCount string `json:"ratingCount" bson:"rating_count,omitempty"`
	RatingValue string `json:"ratingValue" bson:"rating_value,omitempty"`
}
