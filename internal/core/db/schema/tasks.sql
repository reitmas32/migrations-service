CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Necesario para gen_random_uuid()

CREATE TABLE tasks (
                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Cambia a UUID con generación automática
                    user_id UUID NOT NULL, -- Cambia a UUID
                    task_id VARCHAR(40) NOT NULL,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);