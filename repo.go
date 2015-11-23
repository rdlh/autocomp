package main

import (
	"math/rand"
)

var currentId int
var documents Documents
var letters = []rune("abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ123456789")

// Give us some seed data
func init() {
}

func RepoGetDocuments(query string, apiKey string) []string {
	var replies []string

	c, _ := redisPool.Get()
  replies, _ = c.Cmd("ZRANGEBYLEX", apiKey + ":autocomplete", "[" + query, "[" + query + "\xff").List()

  return replies
}

func RepoCreateDocument(d Document) Document {
	currentId += 1
	d.Id = currentId
	documents = append(documents, d)
	return d
}

// func RepoGetAuthKey(a AuthKey) (string, error) {
// 	reply := client.Cmd("HGET", "auth_keys", key)
// 	if reply.Type == redis.NilReply {
// 		return "", reply.Err
// 	}
// 	domain, err := reply.Str()
// 	if err != nil {
// 		return "", err
// 	} else {
// 		return domain, nil
// 	}
// }

func RepoCreateAuthKey(owner string) AuthKey {
	c, _ := redisPool.Get()
	apiKey := randomApiKey(10)
	c.Cmd("HSET", "authkeys", apiKey, owner)
	return AuthKey{ApiKey: apiKey, Owner: owner}
}

func randomApiKey(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
