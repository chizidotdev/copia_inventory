package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queries interface {
	CreateUser(_ context.Context, arg CreateUserParams) (User, error)
	UpsertUser(_ context.Context, arg CreateUserParams) (User, error)
	GetUserByEmail(_ context.Context, email string) (User, error)

	ListOrders(ctx context.Context, userID uuid.UUID) ([]Order, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	DeleteOrder(ctx context.Context, arg DeleteOrderParams) error
	GetOrder(ctx context.Context, id uuid.UUID) (Order, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)

	//ListOrderItems(ctx context.Context, orderID uuid.UUID) ([]OrderItem, error)
	//CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error)
	//DeleteOrderItem(ctx context.Context, id uuid.UUID) error
	//GetOrderItem(ctx context.Context, id uuid.UUID) (OrderItem, error)
	//UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error)

	// GetReports(ctx context.Context, userEmail string) (datastruct.ReportRow, error)
}

var _ Queries = (*Repository)(nil)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	err := db.AutoMigrate(&User{}, &OrderItem{}, &Order{})
	if err != nil {
		log.Panic("Cannot migrate db:", err)
	}

	return &Repository{
		DB: db,
	}
}

func (r *Repository) WithTx(tx *gorm.DB) *Repository {
	return &Repository{
		DB: tx,
	}
}

// ExecTx executes a function within a transaction.
func (r *Repository) ExecTx(ctx context.Context, fn func(*Repository) error) error {
	tx := r.DB.Begin()

	qtx := r.WithTx(tx)

	err := fn(qtx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit().Error
}
