-- creates an enum type for status
create type account_status_type as enum ('new', 'verified', 'suspended', 'inactive');

-- creates an enum for types of accounts
create type account_type as enum ('user', 'org');

-- this table contains the all accounts of this system.
create table if not exists accounts (
    -- account ID
    id                      uuid not null,
    -- account's username
    username                citext unique not null,
    -- account's email address
    email                   citext unique not null,
    -- account's password
    password                text not null,
    -- account's name
    name                    text not null,
    -- status can be one of the following: new, verified, suspended, inactive
    status                  account_status_type not null default 'new',
    -- gravatar url for profile picture
    picture_url             text not null,
    -- type can be either "user" or "org"
    type                    account_type not null default 'user',
    -- timestamp of when the account was created
    created_at              timestamptz not null default current_timestamp,
    -- verification code expected to verify this account
    verification_code       uuid not null,
    -- expiration time for the verification code
    code_expiration         timestamptz not null,
    -- timestamp of when the account was verified
    verified_at             timestamptz not null,
    -- timestamp of when the account was last updated
    updated_at              timestamptz not null,

    primary key(id)
);
