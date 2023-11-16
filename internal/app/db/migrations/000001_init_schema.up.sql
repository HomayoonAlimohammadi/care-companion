create table if not exists care_seeker (
    id bigserial primary key,
    name varchar(256) not null
);

create table if not exists country (
    id bigserial primary key,
    name varchar(256) not null
);

create table if not exists city (
    id bigserial primary key,
    name varchar(256) not null,
    country_id bigint not null references country(id)
);

create table if not exists neighborhood (
    id bigserial primary key,
    name varchar(256) not null,
    city_id bigint not null references city(id)
);

create table if not exists location (
    id bigserial primary key,
    country_id bigint not null references country(id),
    city_id bigint not null references city(id),
    neighborhood_id bigint references neighborhood(id),
    latitude decimal(9, 6),
    longitude decimal(9, 6)
);

create table if not exists care_seek (
    id bigserial primary key,
    location_id bigint not null references location(id),
    compensation decimal(10, 2) not null,
    care_seeker_id bigint not null references care_seeker(id)
);
