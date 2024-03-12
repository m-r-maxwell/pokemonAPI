\connect pokemon
create schema moves;
create table moves.moves (
    move_id serial primary key,
    move_name varchar(100) not null,
    move_type varchar(100) not null,
    move_category varchar(100) not null,
    move_power integer not null,
    move_accuracy integer not null,
    move_pp integer not null,
    move_effect text not null,
    move_effect_chance double not null,
    move_status_condition varchar(100) not null,
    move_priority integer
);
