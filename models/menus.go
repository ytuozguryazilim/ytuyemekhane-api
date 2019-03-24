package models

import (
	"errors"
)

type Menus []Menu

func (ms Menus) SearchMenuByDate(d Date) (Menus, error) {
	var result Menus
	for _, m := range ms {
		m.Print()
		if m.Date.ToString() == d.ToString() {
			result = append(result, m)
			return result, nil
		}
	}
	return result, errors.New("Bu tarihe ait herhangi bir menu yok")
}

func (ms Menus) Print() {
	for _, m := range ms {
		m.Print()
	}
}
