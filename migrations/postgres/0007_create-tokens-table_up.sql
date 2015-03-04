-- creates an enum type for token statuses
create type token_status_type as enum ('new', 'expired', 'revoked');

-- creates an enum type for OAuth2 token types
create type oauth_token_type as enum ('access', 'refresh');

-- this table contains all oauth2 tokens generated for a client and
-- grant code, upon resource owner approval to consume her resources with a given
-- grant's scope.
create table if not exists tokens (
	-- Token value.
	token               uuid not null,
	-- Token type, for now only "bearer"
	token_type          citext not null,
	-- OAuth2 token type, either access or refresh token.
	oauth_type          oauth_token_type not null default 'access',
	-- client ID to which this token is being generated.
	client_id           uuid not null references clients (id),
	-- Grant used to create the token for the first time.
	grant_code          uuid not null references grants (code),
	-- Embedded token scope. The first time it will be the same grant's scope.
	-- However, it could get shrink by the client upon eventual token refreshes,
	-- but certainly not expanded any further.
	scope               citext not null,
	-- status can be one of the following: new, expired, revoked.
	status              token_status_type not null default 'new',
	-- timestamp of when the token was created.
	created_at          timestamptz not null default current_timestamp,
	-- timestamp of when the token was revoked.
	revoked_at          timestamptz not null,
	 -- expiration time for this token.
	expires_in          timestamptz not null,

	primary key (token)
);
