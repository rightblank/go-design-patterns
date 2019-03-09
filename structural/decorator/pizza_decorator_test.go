package decorator

import (
    "strings"
    "testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
    pizza := &PizzaDecorator{}
    pizzaResult, _ := pizza.AddIngredient()
    expectedText := "Pizza with the following ingredients:"
    if !strings.Contains(pizzaResult, expectedText) {
        t.Errorf("Expected result: %s\n not: %s\n", pizzaResult, expectedText)
    }
}

func TestOnion_AddIngredient(t *testing.T) {
    onion := &Onion{}
    onionResult, err := onion.AddIngredient()
    if err == nil {
        t.Errorf("Calling Onnion decorator without an IngredientAdd on its Ingredient field,\n " +
            "must return an error , not a string with '%s'\n", onionResult)
    }

    onion = &Onion{&PizzaDecorator{}}
    onionResult, err = onion.AddIngredient()
    if err != nil {
        t.Error(err)
    }
    if !strings.Contains(onionResult, "onion") {
        t.Errorf("When calling the add ingredient of the onion decorator it must reuturn a text with the word 'onion'," +
            " not '%s", onionResult)
    }
}
