package models

// A vendor selling an article
type Vendor = string

// The id of an article
type ArticleID = int

// A struct representing a purchase of an article from a vendor
type Purchase struct {
	Vendor Vendor `json:"lieferant,omitempty"`

	ArticleID ArticleID `json:"articleID,omitempty"`

	Bulk int `json:"menge,omitempty"`

	Price float64 `json:"preis,omitempty"`
}

func (p Purchase) IsValid() bool {
	return p.ArticleID > 0 && len(p.Vendor) > 0 && p.Bulk > 0 && p.Price >= 0
}
