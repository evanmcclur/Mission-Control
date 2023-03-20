-- +migrate Up
-- +migrate StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.articles (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title varchar NOT NULL,
    strapline varchar,
    published timestamp,
    link varchar NOT NULL,
    author varchar,
    archived boolean NOT NULL DEFAULT FALSE
)

-- +migrate StatementEnd


-- +migrate Down
DROP TABLE articles;