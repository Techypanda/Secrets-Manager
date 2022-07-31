CREATE KEYSPACE IF NOT EXISTS "SecretsManager" WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};
USE "SecretsManager";
CREATE TABLE IF NOT EXISTS secrets(
  user_id int,
  rotation_time timestamp,
  secret_id UUID PRIMARY KEY,
  encrypted_secret blob
);
CREATE TABLE IF NOT EXISTS user(
  user_id int PRIMARY KEY,
  encryption_key blob
);