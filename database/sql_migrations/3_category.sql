-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id SERIAL,
  name VARCHAR(256),
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default NULL,
  PRIMARY KEY (id)
)

-- +migrate StatementEnd


