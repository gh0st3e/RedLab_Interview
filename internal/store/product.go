package store

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	dirName = "files"
)

type ProductStore struct {
}

func NewProductStore() *ProductStore {
	return &ProductStore{}
}

// SaveProduct func which allows to save product into dir
func (p *ProductStore) SaveProduct(name, strProduct string) error {
	err := os.Chdir(dirName)
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strProduct)

	return err

}

// RetrieveProduct func which allows to get product from dir using file name
func (p *ProductStore) RetrieveProduct(name string) (string, error) {
	err := os.Chdir(dirName)
	if err != nil {
		return "", err
	}
	defer os.Chdir("..")

	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data := make([]byte, 128)

	for {
		_, err := f.Read(data)
		if err == io.EOF {
			break
		}
	}

	return string(data), err
}

// DeleteProduct func which allows to delete product from dir
func (p *ProductStore) DeleteProduct(name string) error {
	err := os.Chdir(dirName)
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = os.Remove(name)

	return err
}

// RetrieveProductsByUserID func which allows to get every user's product
func (p *ProductStore) RetrieveProductsByUserID(userID string) ([]string, error) {
	dir, err := os.Open(dirName)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, file := range files {
		if strings.Contains(file.Name(), userID) {
			strProduct, err := p.RetrieveProduct(file.Name())
			fmt.Println(file.Name())
			if err != nil {
				return nil, err
			}
			result = append(result, strProduct)
		}
	}

	return result, nil
}
