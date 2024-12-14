CREATE TABLE IF NOT EXISTS properties (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    location text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    created_by text NOT NULL,
    version integer NOT NULL DEFAULT 1
);