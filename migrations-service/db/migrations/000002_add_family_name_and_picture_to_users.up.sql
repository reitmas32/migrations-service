-- Agregar columnas family_name y picture a la tabla users
ALTER TABLE users
    ADD COLUMN family_name VARCHAR(255),
ADD COLUMN picture TEXT;