\c www
create schema if not exists blog;

create table if not exists blog.thumbnails
(
    id         serial primary key not null,
    alt_text   varchar(256)       not null,
    image      varchar(512)       not null,
    width      integer            not null,
    height     integer            not null,
    type       varchar(32)        not null default 'image/png'
);


create table if not exists blog.posts
(
    id           int primary key not null,
    title        varchar(60)     not null,
    description  varchar(180)    not null,
    created_at   timestamp       not null default current_timestamp,
    content      text            not null,
    lang         varchar(2)      not null default 'en',
    thumbnail_id int             not null references blog.thumbnails (id)
);
