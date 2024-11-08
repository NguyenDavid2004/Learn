-- +goose Up
-- +goose StatementBegin
CREATE TABLE classes (
     id SERIAL PRIMARY KEY,
     name VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd
