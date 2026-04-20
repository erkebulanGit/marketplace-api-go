CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    category_id INTEGER NOT NULL REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );