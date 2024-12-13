-- Eliminar columnas family_name y picture de la tabla users
ALTER TABLE users
DROP COLUMN family_name,
DROP COLUMN picture;