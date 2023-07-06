CREATE TABLE IF NOT EXISTS users
(
    id serial INT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name TEXT,
    email TEXT,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
