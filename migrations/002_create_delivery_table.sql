-- +goose Up
CREATE TABLE delivery (
    order_uid VARCHAR REFERENCES orders(order_uid) ON DELETE CASCADE,
    name VARCHAR,
    phone VARCHAR,
    zip VARCHAR,
    city VARCHAR,
    address VARCHAR,
    region VARCHAR,
    email VARCHAR
);

-- +goose Down
DROP TABLE delivery;
