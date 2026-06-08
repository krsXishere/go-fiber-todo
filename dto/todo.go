package dto

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateTodo struct {
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
}

type UpdateTodo struct {
	ID       string `json:"-"`
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
}
