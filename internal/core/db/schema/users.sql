CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Requerido para gen_random_uuid()

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                       family_name VARCHAR(255),
                       picture TEXT
);