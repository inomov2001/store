CREATE TABLE IF NOT EXISTS fructs (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR,
    price FLOAT,
    quantity FLOAT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS drinks (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR,
    brand VARCHAR,
    price FLOAT,
    liter FLOAT,
    count BIGINT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP
);