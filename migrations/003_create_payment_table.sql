-- +goose Up
CREATE TABLE payment (
    order_uid VARCHAR REFERENCES orders(order_uid) ON DELETE CASCADE,
     transaction VARCHAR,
    request_id VARCHAR,
    currency VARCHAR,
    provider VARCHAR,
    amount INTEGER,
    payment_dt BIGINT,
    bank VARCHAR,
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

-- +goose Down
DROP TABLE payment;
