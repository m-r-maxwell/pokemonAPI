\connect pokemon
create schema items;
create table items.items (
    item_id serial primary key,
    item_name varchar(100) not null,
    item_description text not null,
    item_price integer not null
);
