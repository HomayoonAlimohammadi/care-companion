create table care_seeker (
    id serial primary key,
    name varchar(256) not null
);

create table country (
    id serial primary key,
    name varchar(256)
);

create table city (
    id serial primary key,
    name varchar(256),
    country_id int references country(id)
);

create table neighborhood (
    id serial primary key,
    name varchar(256),
    city_id int references city(id)
);

create table location (
    id serial primary key,
    country_id int references country(id),
    city_id int references city(id),
    neighborhood_id int references neighborhood(id),
    latitude decimal(9, 6),
    longitude decimal(9, 6),
);

create table care_need (
    id serial primary key,
    location_id int references location(id),
    compensation decimal(10, 2)
);

-- TODO: where I left june 13th 23:08 