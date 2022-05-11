package forumDemo

import "github.com/google/uuid"

// Thread
// Des structs, comme en TS qui vont définir la forme de mes entités
type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"ID"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

// ThreadStore
// Des interfaces comme en PHP à la différence près qu'on ne va pas avoir besoin
// de les "attacher" en les déclarant à ma classe comme en PHP, ici, c'est plus libre
// Si une entité possède toutes les méthodes, alors elle implémente l'interface
// C'est plus un nom pour un truc qui implémente tout plutôt qu'une contrainte réelle
type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error // Je passe en référence pour pouvoir modifier mon entité dans la méthode
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
