CREATE TABLE IF NOT EXISTS reviews (
                                       id SERIAL PRIMARY KEY,
                                       product_id INTEGER NOT NULL REFERENCES products(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );