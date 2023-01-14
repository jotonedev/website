\c www
create schema if not exists blog;
SET search_path TO blog;

CREATE EXTENSION citext;
CREATE DOMAIN email AS citext
    CHECK ( value ~
            '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

create table if not exists blog.authors
(
    id         serial primary key not null,
    name       varchar(255),
    surname    varchar(255),
    email      citext,
    username   varchar(24)        not null,
    created_at timestamp          not null default current_timestamp
);

CREATE UNIQUE INDEX lower_username_unique ON blog.authors ((lower(username)));

create table if not exists blog.posts
(
    id          int primary key not null check (id between 1000000 and 9999999),
    title       varchar(60)     not null,
    description varchar(180)    not null,
    created_at  timestamp       not null default current_timestamp,
    updated_at  timestamp       not null default current_timestamp,
    author_id   int             not null references blog.authors (id),
    content     text            not null,
    lang        varchar(2)      not null default 'en'
);

CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE OR REPLACE TRIGGER update_post_update
    BEFORE UPDATE
    ON blog.posts
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();
