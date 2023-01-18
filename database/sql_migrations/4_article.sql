-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE articles (
  id SERIAL,
  user_id INT,
  category_id INT,
  title VARCHAR(256),
  description TEXT,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
)

-- +migrate StatementEnd


