CREATE TABLE IF NOT EXISTS group_members (
    id SERIAL PRIMARY KEY,
    group_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    balance DOUBLE PRECISION DEFAULT 0,
    UNIQUE (group_id, user_id)
);
