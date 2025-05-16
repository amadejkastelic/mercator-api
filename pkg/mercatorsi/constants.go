package mercatorsi

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

type SortField string

const (
	SortFieldScore  SortField = "_score"
	SortFieldRating SortField = "rating"
	SortFieldPrice  SortField = "current_price"
	SortFieldWeight SortField = "weight"
	SortFieldNew    SortField = "first_time_online"
)
