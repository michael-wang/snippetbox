package mysql

import (
	"database/sql"

	"github.com/michael-wang/snippetbox/pkg/models"
)

// SnippetModel holds DB and exported methods.
type SnippetModel struct {
	DB *sql.DB
}

// Insert one snippet record to DB.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get returns one Snippet with specified id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest returns 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
