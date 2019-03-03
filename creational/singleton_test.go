package creational

import "testing"

func TestGetInstance(t *testing.T) {
    instance := GetInstance()
    if instance == nil {
        t.Error("instance creation failed\n")
    }

    instance.AddOne()
    if instance.count != 1 {
        t.Error("count not equal to 1\n")
    }

    instance2 := GetInstance()
    instance2.AddOne()
    if instance2.count != 2 {
        t.Error("two instance created\n")
    }
}
