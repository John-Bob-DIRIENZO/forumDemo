package postgres

import (
	"fmt"
	"forumDemo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (forumDemo.Post, error) {
	var p forumDemo.Post
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return forumDemo.Post{}, fmt.Errorf("error getting post: %w", err)
	}

	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]forumDemo.Post, error) {
	var pp []forumDemo.Post
	if err := s.Select(&pp, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []forumDemo.Post{}, fmt.Errorf("error getting posts: %w", err)
	}

	return pp, nil
}

func (s *PostStore) CreatePost(p *forumDemo.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (s *PostStore) UpdatePost(p *forumDemo.Post) error {
	if err := s.Get(p, `UPDATE posts SET thread_id = $2, title = $3, content = $4, votes = $5 WHERE id = $1 RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}

	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	return nil
}
