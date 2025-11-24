create schema if not exists products;

create table if not exists products.product
(
    id SERIAL PRIMARY KEY,
    name_product Varchar(30),
    price Varchar(30),
    id_user INTEGER REFERENCES users.user(id)
)