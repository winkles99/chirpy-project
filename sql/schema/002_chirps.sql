-- 002_chirps.sql

CREATE TABLE chirps ( id UUID PRIMARY KEY DEFAULT gen_random_uuid(), created_at TIMESTAMP NOT NULL DEFAULT NOW(), updated_at TIMESTAMP NOT NULL DEFAULT NOW(), user_id UUID NOT NULL, body TEXT NOT NULL

);