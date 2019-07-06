package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//here we are creating some sample data in practice a database should be used
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticle(id int) (*article, error) {
	for _, a := range articleList {
		if id == a.ID {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
