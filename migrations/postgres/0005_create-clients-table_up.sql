-- this table contains all oauth2 client applications registered in this system.
create table if not exists clients (
    -- client ID.
    id                  uuid not null,
    -- account ID to which the client belongs to.
    account_id          uuid not null references accounts (id),
    -- clients's secret, encrypted.
    secret              text not null,
    -- client's name.
    name                text not null,
    -- client's description.
    description         text not null,
    -- client's app logo url.
    logo_url            text not null,
    -- homepage url for this client.
    homepage_url        text not null,
    -- redirect url to where the authorization server redirects resource owners.
    redirect_url        citext not null,
    -- timestamp of when the client was created.
    created_at          timestamptz not null default current_timestamp,
    -- timestamp of when the client was last updated.
    updated_at          timestamptz not null,

    primary key(id)
);
