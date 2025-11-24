create schema if not exists balance;

create table if not exists balance.balance
(
    id SERIAL PRIMARY KEY,
    money Varchar(20),
    id_user INTEGER REFERENCES users.user(id)
)