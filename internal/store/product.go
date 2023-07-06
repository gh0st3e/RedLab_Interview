package store

import (
	"context"
	"errors"
	"github.com/gh0st3e/RedLab_Interview/internal/entity"
)

// SaveProduct func which allows to save product into dir
func (s *Store) SaveProduct(ctx context.Context, product entity.Product) error {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `INSERT INTO products(barcode,name,description,cost,user_id) VALUES($1,$2,$3,$4,$5)`

	_, err := s.db.ExecContext(ctx, query,
		product.Barcode,
		product.Name,
		product.Desc,
		product.Cost,
		product.UserID)

	if isUniqueViolation(err) {
		return errors.New("product with this barcode already exists")
	}

	return err
}

// RetrieveProduct func which allows to get product from dir using file name
func (s *Store) RetrieveProduct(ctx context.Context, productID string, userID int) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `SELECT barcode,name,description,cost,user_id FROM products WHERE barcode=$1 AND user_id=$2`

	var product entity.Product

	err := s.db.QueryRowContext(ctx, query, productID, userID).Scan(
		&product.Barcode,
		&product.Name,
		&product.Desc,
		&product.Cost,
		&product.UserID)

	return &product, err
}

// DeleteProduct func which allows to delete product from dir
func (s *Store) DeleteProduct(ctx context.Context, productID string, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `DELETE FROM products WHERE barcode=$1 AND user_id=$2`

	res, err := s.db.ExecContext(ctx, query, productID, userID)

	affectedRows, err := res.RowsAffected()
	if affectedRows == 0 {
		return errors.New("no such product to delete")
	}

	return err
}

// RetrieveProductsByUserID func which allows to get every user's product
func (s *Store) RetrieveProductsByUserID(ctx context.Context, userID int) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `SELECT barcode,name,description,cost,user_id FROM products WHERE user_id=$1`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var products []entity.Product

	for rows.Next() {
		product := entity.Product{}

		err := rows.Scan(
			&product.Barcode,
			&product.Name,
			&product.Desc,
			&product.Cost,
			&product.UserID)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, err
}
