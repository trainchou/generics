package alg

import (
	"errors"
	"fmt"
)

// Slice is a type to score elem.
type Slice []interface{}

// NewSlice is a func to new a Slice.
func NewSlice() Slice {
	return make(Slice, 0)
}

// Add is a func to add an elem to a Slice.
func (s *Slice) Add(elem interface{}) error {
	for _, v := range *s {
		if v == elem {
			fmt.Printf("Slice:Add elem: %v already exist\n", elem)
			return errors.New("elem is already exist")
		}
	}
	*s = append(*s, elem)
	fmt.Printf("SLice:Add elem:%v succ\n", elem)
	return nil
}

// Remove is a func to remove an elem from a Slice.
func (s *Slice) Remove(elem interface{}) error {
	found := false
	for i, v := range *s {
		if v == elem {
			if i == len(*s)-1 {
				*s = (*s)[:i]
			} else {
				*s = append((*s)[:i], (*s)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Slice:Remove elem:%v not exist\n", elem)
		return errors.New("elem is not exist")
	}
	fmt.Printf("Slice:Remove elem: %v succ\n", elem)
	return nil
}
