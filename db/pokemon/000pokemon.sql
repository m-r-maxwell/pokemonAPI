\connect pokemon
create schema pokemon;
CREATE TABLE pokemon.pokemon (
    pokedex_number SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type1 VARCHAR(100) NOT NULL,
    type2 VARCHAR(100),
    stats jsonb not null
);

create table pokemon.types (
    type_id serial primary key,
    type_name varchar(100) not null
);

create table pokemon.abilities (
    ability_id serial primary key,
    ability_name varchar(100) not null,
    ability_description text not null
);