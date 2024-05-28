package schema

import "github.com/google/uuid"

var UserSchema = `
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

`

type User struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Username string    `db:"username"`
	Password string    `db:"password"`

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
