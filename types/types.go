package types

import (
	"time"
)

// Interface because it is simpler to test interfaces in go
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}


type User struct {
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"_"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID int     `json:"id"`
	Name string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}
type Order struct {
	ID int `json:"id"`
	UserID int `json:"userID"`
	Total float64 `json:"total"`
	Status string `json:"status"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID int `json:"id"`
	OrderID int `json:"orderID"`
	ProductID int `json:"productID"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}



type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResgisterPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"` 
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity int 	`json:"quantity"`
}

type CartCheckoutPayload struct {
	Items[] CartCheckoutItem `json:"items" validate:"required"`
}