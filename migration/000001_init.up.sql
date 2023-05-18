CREATE TABLE passports 
(
    id serial primary key,
    type VARCHAR(255) not null,
    number VARCHAR(255) not null unique
);

CREATE TABLE departments 
(
    id serial primary key,
    name VARCHAR(255) not null,
    phone VARCHAR(255) not null unique
);

CREATE TABLE companies 
(
    id serial primary key,
    name VARCHAR(255) not null unique
);

CREATE TABLE employees
(
    id serial not null unique,
    name VARCHAR(255) not null,
    surname VARCHAR(255) not null,
    phone VARCHAR(255) not null unique,
    company_id int not null,
    passport_id int not null unique,
    department_id int not null,
    foreign key (company_id) references companies (id),
    foreign key (passport_id) references passports (id) on delete cascade,
    foreign key (department_id) references departments (id)
);