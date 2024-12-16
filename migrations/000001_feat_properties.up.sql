CREATE TABLE IF NOT EXISTS properties (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    location text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    created_by text NOT NULL,
    version integer NOT NULL DEFAULT 1
);

CREATE INDEX IF NOT EXISTS properties_title_idx ON 
properties USING GIN(to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS properties_description_idx ON 
properties USING GIN(to_tsvector('simple',description));
CREATE INDEX IF NOT EXISTS properties_location_idx ON 
properties USING GIN(to_tsvector('simple',location));