CREATE TABLE IF NOT EXISTS products
(
    barcode text NOT NULL UNIQUE,
    name text NOT NULL,
    desc text NOT NULL,
    cost int NOT NULL,
    user_id int,
    CONSTRAINT products_pkey PRIMARY KEY (barcode),
    CONSTRAINT users_fkey FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
)