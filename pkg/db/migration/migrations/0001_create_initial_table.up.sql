CREATE TABLE events (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);