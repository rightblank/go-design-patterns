package prototype

import "testing"

func TestClone(t *testing.T) {

    shirtsCache := GetShirtsCloner()
    if shirtsCache == nil {
        t.Fatal("cloner not implemented yet")
    }

    item1, err := shirtsCache.GetClone(White)
    if err != nil {
        t.Error(err)
    }

    if item1 == whiteShirtPrototype {
        t.Error("item1 cannot be equal to the white prototype")
    }

    shirt1, ok := item1.(*Shirt)
    if !ok {
        t.Fatal("type assertion for shirt1 failed")
    }

    shirt1.SKU = "abbcc"

    item2, err := shirtsCache.GetClone(White)
    if err != nil {
        t.Fatal(err)
    }

    shirt2, ok := item2.(*Shirt)
    if !ok {
        t.Fatal("type assertion for shirt2 failed")
    }

    if shirt1.SKU == shirt2.SKU {
        t.Error("the SKU of shirt1 and shirt2 must be different")
    }
    if shirt1 == shirt2 {
        t.Error("shirt1 cannot be equal to shirt2")
    }
    t.Logf("LOG: %s", item1.GetInfo())
    t.Logf("LOG: %s", item2.GetInfo())

    t.Logf("LOG: memory address of item1 %p, item2 %p\n", item1, item2)
    t.Logf("LOG: memeory address of the shirts pointer\nshirt1 := %p\nshirt2 := %p", shirt1, shirt2)
    t.Logf("LOG: memeory address of the shirts\n&shirt1 := %p\n&shirt2 := %p", &shirt1, &shirt2)
}
