-- escrow.sql

CREATE TABLE escrow_transactions (
    id uuid PRIMARY KEY,
    buyer_id uuid NOT NULL,
    seller_id uuid NOT NULL,
    bitcoin_address VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('pending', 'confirmed', 'disputed', 'released')),
    created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE disputes (
    id uuid PRIMARY KEY,
    escrow_id uuid REFERENCES escrow_transactions(id),
    reason TEXT NOT NULL,
    resolved BOOLEAN DEFAULT FALSE,
    resolution TEXT,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);
