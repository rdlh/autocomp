package main

type Document struct {
  Id         int       `json:"id"`
	CategoryId int       `json:"category_id"`
  Popularity int       `json:"popularity"`
  Name       string    `json:"name"`
	Value      string    `json:"value"`
}

type Documents []Document
