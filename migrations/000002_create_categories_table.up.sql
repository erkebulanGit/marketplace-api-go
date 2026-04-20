CREATE TABLE IF NOT EXISTS categories (
                                          id SERIAL PRIMARY KEY,
                                          name VARCHAR(255) UNIQUE NOT NULL,
                                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          deleted_at TIMESTAMP
);