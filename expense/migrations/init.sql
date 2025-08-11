CREATE TABLE IF NOT EXISTS expenses (
                                        id SERIAL PRIMARY KEY,
                                        group_id BIGINT NOT NULL,
                                        payer_id BIGINT NOT NULL,
                                        amount DOUBLE PRECISION NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);