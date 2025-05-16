package mercatorsi

import "encoding/json"

type Allergen struct {
	Value     string `json:"value,omitempty"`
	HoverText string `json:"hover_text,omitempty"`
}

type Gtin struct {
	Value string `json:"gtin,omitempty"`
}

type ProductData struct {
	Cinv                          string          `json:"cinv,omitempty"`
	Code                          string          `json:"code,omitempty"`
	Codewz                        string          `json:"codewz,omitempty"`
	Name                          string          `json:"name,omitempty"`
	UnitQuantity                  string          `json:"unit_quantity,omitempty"`
	InvoiceUnit                   string          `json:"invoice_unit,omitempty"`
	InvoiceUnitType               string          `json:"invoice_unit_type,omitempty"`
	AverageWeight                 int             `json:"average_weight,omitempty"`
	NormalPrice                   int             `json:"normal_price,omitempty"`
	CurrentPrice                  string          `json:"current_price,omitempty"`
	Pc30Price                     string          `json:"pc30_price,omitempty"`
	PricePerUnit                  float64         `json:"price_per_unit,omitempty"`
	PricePerUnitBase              string          `json:"price_per_unit_base,omitempty"`
	Eko                           string          `json:"eko,omitempty"`
	HasRecipes                    string          `json:"has_recipes,omitempty"`
	BrandName                     string          `json:"brand_name,omitempty"`
	Discounts                     json.RawMessage `json:"discounts,omitempty"`
	RatingsSum                    string          `json:"ratings_sum,omitempty"`
	RatingsNum                    string          `json:"ratings_num,omitempty"`
	Rating                        int             `json:"rating,omitempty"`
	Package                       int             `json:"package,omitempty"`
	OfferExpiresOn                string          `json:"offer_expires_on,omitempty"`
	Category1                     string          `json:"category1,omitempty"`
	Category2                     string          `json:"category2,omitempty"`
	Category3                     string          `json:"category3,omitempty"`
	AnalyticsObject               json.RawMessage `json:"analytics_object,omitempty"`
	PersonalOfferRecommendationID int             `json:"personal_offer_recommendation_id,omitempty"`
	Allergens                     []Allergen      `json:"allergens,omitempty"`
	Gtins                         []Gtin          `json:"gtins,omitempty"`
}

type Meta struct {
	EsScore float64 `json:"es_score,omitempty"`
	McScore string  `json:"mc_score,omitempty"`
}

type Product struct {
	Data         ProductData `json:"data,omitempty"`
	ItemID       string      `json:"itemId,omitempty"`
	ClassName    string      `json:"className,omitempty"`
	Type         string      `json:"type,omitempty"`
	MainImageSrc string      `json:"mainImageSrc,omitempty"`
	Meta         Meta        `json:"_meta,omitempty"`
	Url          string      `json:"url,omitempty"`
	ShortName    string      `json:"short_name,omitempty"`
	Total        int         `json:"total,omitempty"`
	OrdNum       int         `json:"ordNum,omitempty"`
}

type Category struct {
	ID                 string     `json:"id,omitempty"`
	ParentID           string     `json:"parent_id,omitempty"`
	Name               string     `json:"name,omitempty"`
	UseForInitialSort  string     `json:"use_for_initial_sort,omitempty"`
	Level              string     `json:"level,omitempty"`
	Image              string     `json:"image,omitempty"`
	Note               string     `json:"note,omitempty"`
	HasFoodStuff       string     `json:"has_food_stuff,omitempty"`
	UserPiecesLimit    string     `json:"user_pieces_limit,omitempty"`
	UserKgLimit        string     `json:"user_kg_limit,omitempty"`
	CompanyPiecesLimit string     `json:"company_pieces_limit,omitempty"`
	NoOfProducts       string     `json:"no_of_products,omitempty"`
	SeoDescription     string     `json:"seo_description,omitempty"`
	SeoKeywords        string     `json:"seo_keywords,omitempty"`
	CanRemoveOldDevice string     `json:"can_remove_old_device,omitempty"`
	Created            string     `json:"created,omitempty"`
	Modified           string     `json:"modified,omitempty"`
	Lft                string     `json:"lft,omitempty"`
	Rght               string     `json:"rgt,omitempty"`
	IbmID              string     `json:"ibm_id,omitempty"`
	GoldSort           string     `json:"gold_sort,omitempty"`
	NameWithProdCount  string     `json:"name_with_prod_count,omitempty"`
	ImageURL           string     `json:"image_url,omitempty"`
	Path               string     `json:"path,omitempty"`
	Children           []Category `json:"children,omitempty"`
}

type FilterData struct {
	Categories []Category `json:"categories,omitempty"`
}

type Sort struct {
	Field     SortField     `json:"field,omitempty"`
	Direction SortDirection `json:"direction,omitempty"`
}

type SearchRequest struct {
	Limit      int
	Offset     int
	From       int
	Filter     string
	CategoryID string
	Sort       *Sort
}

type SearchResponse struct {
	Products   []Product  `json:"products,omitempty"`
	FilterData FilterData `json:"filterData,omitempty"`
}

type CategoryAttachment struct {
	Category   Category        `json:"category,omitempty"`
	Attachment json.RawMessage `json:"attachment,omitempty"`
}

type CategoriesResponse struct {
	Values []CategoryAttachment `json:"values,omitempty"`
}
