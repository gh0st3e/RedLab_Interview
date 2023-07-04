package store

import (
	"fmt"
	"os"
)

type ProductStore struct {
}

func NewProductStore() *ProductStore {
	return &ProductStore{}
}

// SaveProduct func which allows to save product into dir
func (p *ProductStore) SaveProduct(name, dirname string, data []byte) error {
	err := os.Chdir(dirname)
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = os.WriteFile(name+".json", data, 0644)
	if err != nil {
		return err
	}

	return err

}

// RetrieveProduct func which allows to get product from dir using file name
func (p *ProductStore) RetrieveProduct(name string, dirname string) ([]byte, error) {
	err := os.Chdir(dirname)
	if err != nil {
		return nil, err
	}
	defer os.Chdir("..")

	f, err := os.Open(name)
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

	return data, nil
}

// DeleteProduct func which allows to delete product from dir
func (p *ProductStore) DeleteProduct(name string, dirname string) error {
	err := os.Chdir(dirname)
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = os.Remove(name)

	return err
}

// RetrieveProductsByUserID func which allows to get every user's product
func (p *ProductStore) RetrieveProductsByUserID(dirname string) ([][]byte, error) {
	dir, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	result := make([][]byte, 0)

	for _, file := range files {
		data, err := p.RetrieveProduct(file.Name(), dirname)
		fmt.Println(file.Name())
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}
