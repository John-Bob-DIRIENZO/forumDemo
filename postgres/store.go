package postgres

import (
	"fmt"
	"forumDemo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Je déclare mon driver comme un side effet pour qu'il soit utilisé par sqlx
)

func NewStore(dataSourceName string) (*Store, error) {
	// Ici je passe un driver qui s'appelle "postgres"... mais ce driver ne veut rien dire en soi
	// il faut que je le télécharge, c'est une librairie
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// On fait un ping pour savoir si elle répond bien
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connectig database: %w", err)
	}

	// Et paf, l'injection de dépendance !
	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil
}

type Store struct {
	forumDemo.ThreadStore
	forumDemo.PostStore
	forumDemo.CommentStore
}
