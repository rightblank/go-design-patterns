package proxy

import (
    "math/rand"
    "testing"
)

func TestUserListProxy_FindUser(t *testing.T) {
    mockedDB := UserList{}

    rand.Seed(3233)
    for i := 0; i < 1000000; i++ {
        n := rand.Int31()
        mockedDB = append(mockedDB, User{ID: n})

    }

    proxy := UserListProxy{
        MockedDatabase: &mockedDB,
        StackSize:      2,
        StackCache:     UserList{},
    }

    knownIDs := [3]int32{
        mockedDB[3].ID,
        mockedDB[4].ID,
        mockedDB[6].ID,
    }
    t.Run("FindUser - Empty cache", func(t *testing.T) {
        user, err := proxy.FindUser(knownIDs[0])
        if err != nil {
            t.Fatal(err)
        }
        if user.ID != knownIDs[0] {
            t.Error("returned id does't match")
        }
        if len(proxy.StackCache) != 1 {
            t.Error("after one complete search, the stack size is not 1")
        }

        if proxy.LastSearchUsedCache == true {
            t.Error(" no user can be returned from an empty cache")
        }
    })

    t.Run("FindUser - One User", func(t *testing.T) {
        user, err := proxy.FindUser(knownIDs[0])
        if err != nil {
            t.Fatal(err)
        }
        if user.ID != knownIDs[0] {
            t.Error("returned id does't match")
        }
        if len(proxy.StackCache) != 1 {
            t.Error("after one complete search, the stack size is not 1")
        }

        if !proxy.LastSearchUsedCache {
            t.Error(" no user can be returned from an empty cache")
        }
    })

    t.Run("FindUser - overflowing the stack", func(t *testing.T) {
        user, err := proxy.FindUser(knownIDs[0])
        if err != nil {
            t.Fatal(err)
        }

        user1, err := proxy.FindUser(knownIDs[1])
        if proxy.LastSearchUsedCache {
            t.Error("user not in cache yet")
        }

        user2, err := proxy.FindUser(knownIDs[2])
        if proxy.LastSearchUsedCache {
            t.Error("user not in cache yet")
        }

        if len(proxy.StackCache) != 2 {
            t.Error("after 3 complete search, the stack size is not 2")
        }

        if user.ID != knownIDs[0] {
            t.Error("returned id does't match")
        }
        if len(proxy.StackCache) != 2 {
            t.Error("after one complete search, the stack size is not 2")
        }

        for _, v := range proxy.StackCache {
            if v != user1 && v != user2 {
                t.Error("unexpected user found in the cache")
            }
        }
    })
}
