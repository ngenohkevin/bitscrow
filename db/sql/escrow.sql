-- escrow.sql

CREATE TABLE escrow_transactions (
    id SERIAL PRIMARY KEY,
    buyer_id INTEGER NOT NULL,
    seller_id INTEGER NOT NULL,
    bitcoin_address VARCHAR(255) NOT NULL,
    amount NUMERIC(18,8) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('pending', 'confirmed', 'disputed', 'released')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE disputes (
    id SERIAL PRIMARY KEY,
    escrow_id INTEGER REFERENCES escrow_transactions(id),
    reason TEXT NOT NULL,
    resolved BOOLEAN DEFAULT FALSE,
    resolution TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
