-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id SERIAL,
  name VARCHAR(256),
  email VARCHAR(256) UNIQUE,
  password VARCHAR(256),
  bio VARCHAR(256),
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default NULL,
  PRIMARY KEY (id)
)

-- +migrate StatementEnd


