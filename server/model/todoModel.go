package model

type Todo struct {
	ID          string `jvalidate:"required" json:"id"`
	Title       string `jvalidate:"required" json:"title"`
	Description string `jvalidate:"required" json:"description"`
	Completed   bool   `jvalidate:"required" json:"completed"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
