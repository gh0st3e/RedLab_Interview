package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gh0st3e/RedLab_Interview/internal/entity"
)

const (
	FileStorage    = "files"
	JSONExtension  = ".json"
	FilePermission = 0644
	DirPermission  = 0777
)

type ProductStore struct {
}

func NewProductStore() (*ProductStore, error) {
	_, err := os.Open(FileStorage)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(FileStorage, DirPermission)
		if err != nil {
			return nil, fmt.Errorf("error while creating file storage: %s", err.Error())
		}
	}
	return &ProductStore{}, nil
}

// SaveProduct func which allows to save product into dir
func (p *ProductStore) SaveProduct(product entity.Product) error {
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	strID := strconv.Itoa(product.UserID)

	err = os.WriteFile(filepath.Join(FileStorage, strID, product.Barcode)+JSONExtension, data, FilePermission)
	if err != nil {
		return err
	}

	return err
}

// RetrieveProduct func which allows to get product from dir using file name
func (p *ProductStore) RetrieveProduct(fileName string, userID int) (*entity.Product, error) {
	strID := strconv.Itoa(userID)
	f, err := os.Open(filepath.Join(FileStorage, strID, fileName+JSONExtension))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, fi.Size())
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}

	product := &entity.Product{}

	err = json.Unmarshal(data, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct func which allows to delete product from dir
func (p *ProductStore) DeleteProduct(fileName string, userID int) error {
	strID := strconv.Itoa(userID)
	return os.Remove(filepath.Join(FileStorage, strID, fileName) + JSONExtension)
}

// RetrieveProductsByUserID func which allows to get every user's product
func (p *ProductStore) RetrieveProductsByUserID(userID int) ([]entity.Product, error) {
	strID := strconv.Itoa(userID)
	dir, err := os.Open(filepath.Join(FileStorage, strID))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	products := make([]entity.Product, len(files))

	for i, file := range files {
		product, err := p.RetrieveProduct(strings.Trim(file.Name(), JSONExtension), userID)
		if err != nil {
			return nil, err
		}
		products[i] = *product
	}

	return products, nil
}

func (p *ProductStore) CreateUserStorage(userID int) error {
	strID := strconv.Itoa(userID)

	return os.Mkdir(filepath.Join(FileStorage, strID), DirPermission)
}
