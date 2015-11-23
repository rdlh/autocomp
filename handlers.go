package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
	"github.com/fzzy/radix/redis"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "Pong\n")
}

func DocumentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	apiKey, success := CheckAuthKey(w, r)
	if success {
		w.WriteHeader(http.StatusOK)
		documents := RepoGetDocuments("a", apiKey)
		if err := json.NewEncoder(w).Encode(documents); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		output := map[string]string{"error": "Api-Key not recognized"}
		if err := json.NewEncoder(w).Encode(output); err != nil {
			panic(err)
		}
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Document"}' http://localhost:8080/documents

*/
func DocumentCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	apiKey, _ := CheckAuthKey(w, r)
	var model Document
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &model); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(apiKey); err != nil {
			panic(err)
		}
	}

	t := RepoCreateDocument(model)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func AuthKeyCreate(w http.ResponseWriter, r *http.Request) {
	authKey := RepoCreateAuthKey(strings.Trim(r.FormValue("owner"), " "))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(authKey); err != nil {
		panic(err)
	}
}

func CheckAuthKey(w http.ResponseWriter, r *http.Request) (string, bool) {
	c, _ := redisPool.Get()
	apiKey := r.Header.Get("X-API-KEY")
	reply := c.Cmd("HGET", "authkeys", apiKey)
	if reply.Type == redis.NilReply {
		log.Printf("FALSE")
		return "", false
	}

	return apiKey, true
}
