CREATE TABLE countries (
    id uuid primary key not null,
    name varchar(40) UNIQUE,
    currency varchar(10),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE cities (
    id uuid primary key not null,
    name varchar(40) UNIQUE,
    population int default 0,
    country_id uuid references countries(id),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

