-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links (
    id VARCHAR PRIMARY KEY,
    user_id uuid NOT NULL,
    short_url VARCHAR NOT NULL,
    long_url VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT now () NOT NULL,
    last_modified TIMESTAMP DEFAULT now () NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;

-- +goose StatementEnd
