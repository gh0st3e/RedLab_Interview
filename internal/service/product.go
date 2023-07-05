package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gh0st3e/RedLab_Interview/internal/entity"

	"github.com/sirupsen/logrus"
)

const (
	JSONExtension = ".json"
)

type ProductActions interface {
	SaveProduct(fileName, userDir string, product entity.Product) error
	RetrieveProduct(fileName string, userDir string) (*entity.Product, error)
	DeleteProduct(fileName string, userDir string) error
	RetrieveProductsByUserID(userDir string) ([]entity.Product, error)
}

type ProductService struct {
	logger       *logrus.Logger
	productStore ProductActions
}

func NewProductService(logger *logrus.Logger, productStore ProductActions) *ProductService {
	return &ProductService{
		logger:       logger,
		productStore: productStore,
	}
}

func (p *ProductService) SaveProduct(userID int, product entity.Product) error {
	p.logger.Info("[SaveProduct] started")

	product.UserID = userID
	strID := strconv.Itoa(userID)

	exProduct, err := p.productStore.RetrieveProduct(product.Barcode+JSONExtension, strID)
	if exProduct != nil {
		p.logger.Info("[SaveProduct] product with this barcode already exists")
		return fmt.Errorf("product with this barcode already exists")
	}

	err = p.productStore.SaveProduct(product.Barcode, strID, product)
	if err != nil {
		p.logger.Errorf("[SaveProduct] error while saving product: %s", err.Error())
		return fmt.Errorf("error while saving product, try later\n%w", err)
	}

	p.logger.Info("[SaveProduct] ended")

	return nil
}

func (p *ProductService) RetrieveProduct(fileName string, userID int) (*entity.Product, error) {
	p.logger.Info("[RetrieveProduct] started")

	strID := strconv.Itoa(userID)

	product, err := p.productStore.RetrieveProduct(fileName+JSONExtension, strID)
	if err != nil {
		p.logger.Errorf("[RetrieveProduct] error while retrieving product: %s", err.Error())
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("no such product")
		}
		return nil, fmt.Errorf("error while retrieveing product\n%w", err)
	}

	p.logger.Info(product)
	p.logger.Info("[RetrieveProduct] ended")

	return product, nil

}

func (p *ProductService) DeleteProduct(fileName string, userID int) error {
	p.logger.Info("[DeleteProduct] started")

	strID := strconv.Itoa(userID)

	err := p.productStore.DeleteProduct(fileName, strID)
	if err != nil {
		p.logger.Errorf("[DeleteProduct] error while deleting: %s", err)
		return fmt.Errorf("error while deleting\n%w", err)
	}

	p.logger.Info("[DeleteProduct] ended")

	return nil
}

func (p *ProductService) RetrieveProductsByUserID(userID int) ([]entity.Product, error) {
	p.logger.Info("[RetrieveProductsByUserID] started")

	strID := strconv.Itoa(userID)

	products, err := p.productStore.RetrieveProductsByUserID(strID)
	if err != nil {
		p.logger.Errorf("[RetrieveProductsByUserID] error while retrieving products: %s", err.Error())
		return nil, err
	}

	p.logger.Info(products)
	p.logger.Info("[RetrieveProductsByUserID] ended")

	return products, nil
}
