-- +migrate Up
CREATE TABLE articles (id int, title varchar);


-- +migrate Down
DROP TABLE articles;