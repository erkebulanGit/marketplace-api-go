CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     email VARCHAR(255) UNIQUE NOT NULL,
                                     phone VARCHAR(50),
                                     password VARCHAR(255) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP
);