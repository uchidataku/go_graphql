CREATE TABLE IF NOT EXISTS todos(
    todo_id serial PRIMARY KEY,
    content VARCHAR (50) UNIQUE NOT NULL
);