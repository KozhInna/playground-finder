CREATE TABLE playgrounds(
    id serial primary key,
    hri_id integer unique not null, 
    name_fi text, 
    name_en text, 
    addr_street text, 
    addr_city text, 
    latitude DOUBLE precision, longitude double precision, 
    www_fi text, 
    phone text, 
    has_building boolean not null, 
    picture_url text, 
    desc_fi text, desc_en text, 
    status text not null default 'active' check (status in ('active', 'removed')), 
    created_at timestamptz not null default now(), 
    last_synced_at timestamptz not null default now()
);

create table users (
    id serial primary key, 
    google_sub text unique not null, 
    registered_at timestamptz not null default now(), 
    name_nick text
);

create table photos (
    id serial primary key, 
    playground_id integer not null references playgrounds(id), 
    photo_url text not null, 
    created_at timestamptz not null default now(), 
    added_by integer references users(id) on delete set null
);
CREATE INDEX ON PHOTOS (playground_id);

create table reviews (
    id serial primary key, 
    playground_id integer not null references playgrounds(id), 
    added_by integer references users(id) on delete set null, 
    description text not null, 
    created_at timestamptz not null default now(), 
    updated_at timestamptz not null default now()
);
CREATE INDEX ON reviews (playground_id);
