CREATE TABLE IF NOT EXISTS accounts(
    account_id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL
);