package main

type AuthKey struct {
  ApiKey string `json:"api_key"`
  Owner  string `json:"owner"`
}

type AuthKeys []AuthKey
