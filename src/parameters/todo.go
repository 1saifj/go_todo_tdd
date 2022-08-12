package parameters

const (
	defaultLimit  int    = 20
	defaultOffset int    = 0
	defaultOrder  string = "created_at desc"
)

type FilterTodoParameters struct {
	Title      *string `url:"title"`
	Limit      *int    `url:"limit"`
	Offset     *int    `url:"offset"`
	IsComplete *bool   `url:"is_complete"`
	OrderBy    *string `url:"order_by"`
	OrderField *string `url:"order_field"`
}

func (filter *FilterTodoParameters) GetLimit() int {
	if filter.Limit == nil {
		return defaultLimit
	}
	return *filter.Limit
}

func (filter *FilterTodoParameters) GetOffset() int {
	if filter.Offset == nil {
		return defaultOffset
	}
	return *filter.Offset
}

func (filter *FilterTodoParameters) GetOrderBy() string {
	if filter.OrderBy == nil {
		return defaultOrder
	}
	return *filter.OrderBy
}
