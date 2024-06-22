
CREATE TYPE order_type AS ENUM ('self_pickup', 'delivery');
CREATE TYPE payment_type AS ENUM ('uzum', 'cash', 'terminal');
CREATE TYPE order_status AS ENUM ('waiting_for_payment', 'collecting', 'delivery', 'waiting_on_branch', 'finished', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    external_id VARCHAR(100) NOT NULL,
    type order_type NOT NULL,
    customer_phone VARCHAR(100),
    customer_name VARCHAR(100),
    customer_id UUID,
    payment_type payment_type NOT NULL,
    status order_status NOT NULL,
    to_address VARCHAR,
    to_location POLYGON,
    discount_amount FLOAT,
    amount FLOAT,
    delivery_price FLOAT,
    paid BOOLEAN DEFAULT FALSE,
    courier_id UUID,
    courier_phone VARCHAR,
    courier_name VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0,
    CONSTRAINT phone_deleted_at UNIQUE (customer_phone, deleted_at),
    CONSTRAINT external_id_deleted_at UNIQUE (external_id, deleted_at)
);

CREATE TABLE IF NOT EXISTS order_products (
    id UUID PRIMARY KEY,
    product_id UUID,
    count INTEGER,
    discount_price FLOAT,
    price FLOAT,
    order_id UUID REFERENCES orders(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS order_status_notes (
    id UUID PRIMARY KEY ,
    order_id UUID REFERENCES orders(id),
    status order_status NOT NULL,
    user_id UUID,
    reason VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);