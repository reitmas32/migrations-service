CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Requerido para gen_random_uuid()

CREATE TABLE devices (
                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                    user_id UUID NOT NULL, -- Cambia a UUID
                    operative_system TEXT NOT NULL,
                    os_version TEXT NOT NULL,
                    token TEXT NOT NULL,
                    model TEXT NOT NULL,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);