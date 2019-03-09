package decorator

import (
    "errors"
    "fmt"
)

type IngredientAdd interface {
    AddIngredient() (string, error)
}

type PizzaDecorator struct {
    Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
    return "Pizza with the following ingredients:", nil
}

type Meat struct {
    Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
    if m.Ingredient == nil {
        return "", errors.New("meat's Inner IngredientAdd field is empty")
    }
    ingredients, err := m.Ingredient.AddIngredient()
    if err != nil {
        return "Error on adding meat onto pizza", err
    }
    return ingredients + " meat", nil

}

type Onion struct {
    Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
    if o.Ingredient == nil {
        return "", errors.New("onion's Inner IngredientAdd field is empty")
    }

    ingredients, err := o.Ingredient.AddIngredient()
    if err != nil {
        return "Error on adding onion onto pizza", err
    }
    return fmt.Sprintf("%s %s", ingredients, "onion"), nil
}