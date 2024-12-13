-- Crear la tabla de usuarios
CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Requerido para gen_random_uuid()

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Crear la tabla de autenticación de proveedores
CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Necesario para gen_random_uuid()

CREATE TABLE user_auth_providers (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Cambia a UUID con generación automática
                                     user_id UUID NOT NULL, -- Cambia a UUID
                                     provider VARCHAR(50) NOT NULL,
                                     provider_user_id VARCHAR(255) NOT NULL,
                                     access_token TEXT,
                                     refresh_token TEXT,
                                     token_expiry TIMESTAMP,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);