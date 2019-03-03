package prototype

import (
	"errors"
	"fmt"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(color ShirtColor) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtColor byte

const (
	Black ShirtColor = 1
	Blue  ShirtColor = 2
	White ShirtColor = 3
)

type ShirtsCache struct {
}

func (s *ShirtsCache) GetClone(color ShirtColor) (ItemInfoGetter, error) {
	//return nil, errors.New("not implemented yet")
	switch color {
	case White:
		//return whiteShirtPrototype, nil
		newItem := *whiteShirtPrototype
		return &newItem, nil
	case Black:
		newItem := *blackShirtPrototype
		return &newItem, nil
	case Blue:
		newItem := *blueShirtPrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %.2f\n", s.SKU, s.Color, s.Price)
}

var whiteShirtPrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackShirtPrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Black,
}

var blueShirtPrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Blue,
}