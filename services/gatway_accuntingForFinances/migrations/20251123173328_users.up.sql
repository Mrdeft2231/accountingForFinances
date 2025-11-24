create schema if not exists users;

create table if not exists users.user
(
    id SERIAL PRIMARY KEY,
    login Varchar(20) UNIQUE not null,
    password_hash Varchar(30) not null,
    email Varchar(40) UNIQUE not null,
    phone Varchar(20) UNIQUE not null
)