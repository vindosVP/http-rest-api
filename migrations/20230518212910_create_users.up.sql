CREATE TABLE users(
                      uuid uuid not null primary key default uuid_generate_v4(),
                      email varchar not null unique,
                      encrypted_password varchar not null
);
