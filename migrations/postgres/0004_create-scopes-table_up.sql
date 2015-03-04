-- this table contains the all scopes of this system.
create table if not exists scopes (
    -- scope ID
    id                  citext not null,
    -- scope's description
    description         text not null,
    -- timestamp of when the scope was created
    created_at          timestamptz not null default current_timestamp,
    -- timestamp of when the scope was last updated
    updated_at          timestamptz not null,

    primary key(id)
);
