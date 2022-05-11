package postgres

import (
	"fmt"
	"forumDemo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// NewThreadStore ici, c'est un peu mon constructeur, techniquement j'utilise l'injection de dépendance en fait.
// Je dis que j'aurais besoin d'une instance de DB, mais je ne sais pas encore comment je vais la récupérer
func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		db, // Notez que je peux ne pas lui donner de nom explicitement du coup
	}
}

// ThreadStore
// Encore une fois, je ne marque pas le nom de mon élément passé par un pointeur
// ça va me permettre d'intégrer le type, mais surtout, toutes les méthodes de mon
// entité DB. Il faut le lire comme class ThreadStore extends sqlx\DB en PHP, mais où
// l'héritage est passé par injection de dépendance, ouais c'est étrange, mais un peu cool
type ThreadStore struct {
	*sqlx.DB
}

// Thread
// cmd + n pour accéder aux raccourcis et implémenter toutes les méthodes de mon interface
func (s *ThreadStore) Thread(id uuid.UUID) (forumDemo.Thread, error) {
	var t forumDemo.Thread
	// Requête préparée automatique
	// dans quelle variable (donc passée en référence) je vais coller le résultat
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return forumDemo.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}

	// Si tout se passe bien, je retourne mon Thread et pas d'erreur
	return t, nil
}

func (s *ThreadStore) Threads() ([]forumDemo.Thread, error) {
	var tt []forumDemo.Thread
	if err := s.Select(&tt, `SELECT * FROM threads`); err != nil {
		return []forumDemo.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}

	return tt, nil
}

func (s *ThreadStore) CreateThread(t *forumDemo.Thread) error {
	// Ce statement va me retourner directement l'élément fraichement inséré, merci Postgres
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}

	return nil
}

func (s *ThreadStore) UpdateThread(t *forumDemo.Thread) error {
	if err := s.Get(t, `UPDATE threads SET title = $2, descritpion = $3 WHERE id = $1 RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}

	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	// Exec me retourne soit une erreur, soit le nombre de rangées affectées par la commande,
	// comme je n'en ai rien à foutre de cette info, je la colle dans un placeholder spécial : "_"
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}

	return nil
}
