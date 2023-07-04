package entity

import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	Barcode string
	Name    string
	Desc    string
	Cost    int
	UserID  int
}

func (p *Product) String() string {
	var res string

	res += p.Barcode + "\n"
	res += p.Name + "\n"
	res += p.Desc + "\n"
	res += strconv.Itoa(p.Cost) + "\n"
	res += strconv.Itoa(p.UserID) + "\n"

	return res
}

func (p *Product) ToStruct(strProduct string) error {
	lines := strings.Split(strProduct, "\n")

	Cost, err := strconv.Atoi(lines[3])
	if err != nil {
		return fmt.Errorf("[Product.ToStruct] Parse Cost error: %s", err.Error())
	}
	UserID, err := strconv.Atoi(lines[4])
	if err != nil {
		return fmt.Errorf("[Product.ToStruct] Parse UserID error: %s", err.Error())
	}

	p.Barcode = lines[0]
	p.Name = lines[1]
	p.Desc = lines[2]
	p.Cost = Cost
	p.UserID = UserID

	return nil
}
