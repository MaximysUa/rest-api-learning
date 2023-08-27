package user

type User struct {
	ID           string `json:"id" bson:"_id, omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

// Структура для работы с дто
// TODO освежить знания по поводу data transfer object
type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
