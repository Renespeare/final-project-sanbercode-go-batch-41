-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE user_credentials (
  id SERIAL,
  user_id INT,
  uuid VARCHAR(256) UNIQUE,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)

-- +migrate StatementEnd

