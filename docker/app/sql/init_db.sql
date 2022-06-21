create database library;

create extension if not exists "uuid-ossp";

create table author (
  id uuid not null
    constraint df_author_id default uuid_generate_v4()
    constraint pk_author_id primary key,
  name varchar(50) not null,
  specialty varchar(50) not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);