create table users (
    id uuid,
    firstname character(255),
    lastname character(255),
    age smallint,
    is_married boolean,
    password character(255)
)

create table products (
    id uuid,
    description text,
    tags text[],
    quantity smallint
)

create table orders (
    id uuid
)

create table orders_products (
    order_id uuid,
    product_id uuid,
    price integer,
    FOREIGN KEY(order_id) REFERENCES orders(id)
    FOREIGN KEY(product_id) REFERENCES products(id)
)

create table users_orders (
    order_id uuid,
    user_id uuid,
    FOREIGN KEY(order_id) REFERENCES orders(id)
    FOREIGN KEY(user_id) REFERENCES users(id)
)