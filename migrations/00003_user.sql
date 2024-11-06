-- +goose Up
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   email VARCHAR(100) NOT NULL UNIQUE,
   password VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;