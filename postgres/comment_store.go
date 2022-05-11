package postgres

import (
	"fmt"
	"forumDemo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (forumDemo.Comment, error) {
	var c forumDemo.Comment
	if err := s.Get(&c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return forumDemo.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return c, nil
}

func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]forumDemo.Comment, error) {
	var cc []forumDemo.Comment
	if err := s.Select(&cc, `SELECT * FROM comments WHERE post_id = $1`, postID); err != nil {
		return []forumDemo.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}

	return cc, nil
}

func (s *CommentStore) CreateComment(c *forumDemo.Comment) error {
	if err := s.Get(c, `INSERT INTO comments VALUES ($1, $2, $3, $4) RETURNING *`,
		c.ID,
		c.PostID,
		c.Content,
		c.Votes); err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}

	return nil
}

func (s *CommentStore) UpdateComment(c *forumDemo.Comment) error {
	if err := s.Get(c, `UPDATE comments SET post_id = $2, content = $3, votes = $4 WHERE id = $1 RETURNING *`,
		c.ID,
		c.PostID,
		c.Content,
		c.Votes); err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}

	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	return nil
}
