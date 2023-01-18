-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE comments (
  id SERIAL,
  user_id INT,
  article_id INT,
  description TEXT,
  created_at TIMESTAMP default now(),
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE
)

-- +migrate StatementEnd


