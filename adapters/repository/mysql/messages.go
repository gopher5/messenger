package mysql

import (
	"database/sql"
	"errors"
	"messanger/models"
	tr "messanger/pkg/error_trace"
)

type Messages struct {
	db *sql.DB
}

func NewMessages(db *sql.DB) *Messages {
	return &Messages{db}
}

func (m *Messages) New(message *models.Message) error {
	res, err := m.db.Exec("INSERT INTO messages (chat_id, user_id, value, time) VALUE (?, ?, ?, ?)",
		message.ChatId, message.UserId, message.Text, message.Time)
	if err != nil {
		return tr.Trace(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return tr.Trace(err)
	}
	message.Id = int(id)

	return nil
}

const getMessagesByChatQuery = `
SELECT * FROM messages
WHERE chat_id = ? AND id < ?
ORDER BY time DESC
LIMIT ?;
`

func (m *Messages) GetByChat(chatId int, lastId int, count int) ([]models.Message, error) {
	var messages []models.Message
	rows, err := m.db.Query(getMessagesByChatQuery, chatId, lastId, count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return messages, nil
		}
	}

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.Id, &message.ChatId, &message.UserId, &message.Text, &message.Time); err != nil {
			return nil, tr.Trace(err)
		}
	}
	return messages, nil
}

func (m *Messages) GetById(id int) (*models.Message, error) {
	var message models.Message
	if err := m.db.QueryRow("SELECT * FROM messages WHERE id = ?", id).Scan(
		&message.Id, &message.ChatId, &message.UserId, &message.Text, &message.Time); err != nil {
		return nil, tr.Trace(err)
	}
	return &message, nil
}

func (m *Messages) Update(id int, text string) error {
	if _, err := m.db.Exec("UPDATE messages SET value=? WHERE id = ?", text, id); err != nil {
		return tr.Trace(err)
	}
	return nil
}

func (m *Messages) Delete(id int) error {
	if _, err := m.db.Exec("DELETE FROM messages WHERE id = ?", id); err != nil {
		return tr.Trace(err)
	}
	return nil
}

func (m *Messages) DeleteByChat(chatId int) error {
	if _, err := m.db.Exec("DELETE FROM messages WHERE chat_id = ?", chatId); err != nil {
		return tr.Trace(err)
	}
	return nil
}
