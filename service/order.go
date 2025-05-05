package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/nguyenhoang711/demo-ddd/aggregate"
	"github.com/nguyenhoang711/demo-ddd/domain/customer"
	"github.com/nguyenhoang711/demo-ddd/domain/customer/memory"
	"github.com/nguyenhoang711/demo-ddd/domain/customer/mongo"
	"github.com/nguyenhoang711/demo-ddd/domain/product"
	prodmemory "github.com/nguyenhoang711/demo-ddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the mongo repo, if we needed parameters, such as connection strings they could be inputted here
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	// Add Products and Update Customer

	return price, nil
}
