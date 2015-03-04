-- creates an enum type for status
create type grant_status_type as enum ('new', 'used', 'expired', 'revoked');

-- this table contains all oauth2 grant authorization codes generated for a client
-- upon resource owner successful approval to consume her resources.
create table if not exists grants (
	-- Grant code
	code                uuid not null,
	-- account ID to which the client belongs to.
	client_id           uuid not null references clients (id),
	 -- expiration time for the authorization grant code.
	expires_in          timestamptz not null,
	-- status can be one of the following: new, used, expired, revoked.
	status              grant_status_type not null default 'new',
	-- redirect url for which the authorization grant was created.
	redirect_url        citext not null,
	-- timestamp of when the grant was created.
	created_at          timestamptz not null default current_timestamp,
	-- timestamp of when the grant was used by the client.
	used_at             timestamptz not null,
	-- timestamp of when the grant was expired.
	expired_at          timestamptz not null,
	-- timestamp of when the grant was revoked.
	revoked_at          timestamptz not null,

	primary key (code)
);

-- Keeps the list of scopes associated to an authorization grant code.
create table if not exists grant_scopes (
	grant_code   uuid not null references grants (code),
	scope_id     citext not null references scopes (id)
);
