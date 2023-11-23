CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    number BIGINT,
    balance BIGINT
);