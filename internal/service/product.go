package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gh0st3e/RedLab_Interview/internal/entity"
	"os"
)

type ProductStore interface {
	SaveProduct(ctx context.Context, product entity.Product) error
	RetrieveProduct(ctx context.Context, productID string, userID int) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productID string, userID int) error
	RetrieveProductsByUserID(ctx context.Context, userID int) ([]entity.Product, error)
}

func (s *Service) SaveProduct(ctx context.Context, product entity.Product, userID int) error {
	s.logger.Info("[SaveProduct] started")

	product.UserID = userID

	exProduct, err := s.store.RetrieveProduct(ctx, product.Barcode, userID)
	fmt.Println(exProduct)
	if exProduct.Barcode != "" {
		s.logger.Info("[SaveProduct] product with this barcode already exists")
		return fmt.Errorf("product with this barcode already exists")
	}

	err = s.store.SaveProduct(ctx, product)
	if err != nil {
		s.logger.Errorf("[SaveProduct] error while saving product: %s", err.Error())
		return fmt.Errorf("error while saving product, try later\n%w", err)
	}

	s.logger.Info("[SaveProduct] ended")

	return nil
}

func (s *Service) RetrieveProduct(ctx context.Context, barcode string, userID int) (*entity.Product, error) {
	s.logger.Info("[RetrieveProduct] started")

	product, err := s.store.RetrieveProduct(ctx, barcode, userID)
	if err != nil {
		s.logger.Errorf("[RetrieveProduct] error while retrieving product: %s", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no such product")
		}
		return nil, fmt.Errorf("error while retrieveing product\n%w", err)
	}

	s.logger.Info(product)
	s.logger.Info("[RetrieveProduct] ended")

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, barcode string, userID int) error {
	s.logger.Info("[DeleteProduct] started")

	err := s.store.DeleteProduct(ctx, barcode, userID)
	if err != nil {
		s.logger.Errorf("[DeleteProduct] error while deleting: %s", err)
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("project with this barcode not exists")
		}
		return fmt.Errorf("error while deleting\n%w", err)
	}

	s.logger.Info("[DeleteProduct] ended")

	return nil
}

func (s *Service) RetrieveProductsByUserID(ctx context.Context, userID int) ([]entity.Product, error) {
	s.logger.Info("[RetrieveProductsByUserID] started")

	products, err := s.store.RetrieveProductsByUserID(ctx, userID)
	if err != nil {
		s.logger.Errorf("[RetrieveProductsByUserID] error while retrieving products.sql: %s", err.Error())
		return nil, err
	}

	s.logger.Info(products)
	s.logger.Info("[RetrieveProductsByUserID] ended")

	return products, nil
}
