BEGIN;

CREATE TABLE IF NOT EXISTS user_balances
(
    id serial primary key not null,
    user_id    uuid not null,
    balance    decimal  null default null,
);

CREATE TABLE IF NOT EXISTS reservation
(
    user_id uuid not null,
    service_id uuid not null,
    order_id uuid not null,
    cost decimal,
);

CREATE TABLE IF NOT EXISTS report
(
    user_id uuid not null,
    service_id uuid not null,
    order_id uuid not null,
    cost decimal,
);

COMMIT;