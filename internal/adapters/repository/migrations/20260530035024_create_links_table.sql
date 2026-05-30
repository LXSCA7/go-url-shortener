-- +goose Up
create table links (
   id bigint primary key,
   original_url text not null,
   short_code varchar(15) unique not null,
   created_at timestamp with time zone default current_timestamp,
   visits integer default 0
);

create index idx_short_code on links(short_code);

-- +goose Down
drop table links;
