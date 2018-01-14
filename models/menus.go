package models

import (
	"errors"
	"fmt"
)

// Menus :
type Menus []Menu

// SearchMenuByDate :
func (ms Menus) SearchMenuByDate(d Date) (Menu, error) {
	fmt.Println(d.ToString())
	for _, m := range ms {
		fmt.Println("##########")
		m.Print()
		fmt.Println(d.ToString())
		if m.Date.ToString() == d.ToString() {
			return m, nil
		}
	}
	return Menu{}, errors.New("Bu tarihe ait herhangi bir menu yok")
}
