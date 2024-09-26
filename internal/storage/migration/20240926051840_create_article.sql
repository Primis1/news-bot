-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles {
    id SERIAL PRIMARY KEY,
    source_id INT NOT NULL,
    title VARCHAR(256) NOT NULL,
    link VARCHAR(256) NOT NULL,
    published_at INT NOT NULL,
    summary TEXT NOT NULL, 
    create_at TIMESTAMP NOW(),
    posted_at TIMESTAMP NOW(),
}
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
