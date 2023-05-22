CREATE TABLE companies 
(
    id serial primary key,
    name VARCHAR(255) not null unique
);

CREATE TABLE departments 
(
    id serial primary key,
    name VARCHAR(255) not null,
    phone VARCHAR(255) not null unique,
    companyId int not null,
    foreign key (companyId) references companies (id)
);

CREATE TABLE employees
(
    id serial not null unique,
    name VARCHAR(255) not null,
    surname VARCHAR(255) not null,
    phone VARCHAR(255) not null unique,
    companyId int not null,
    departmentId int not null,
    foreign key (companyId) references companies (id),
    foreign key (departmentId) references departments (id)
);

CREATE TABLE passports 
(
    id serial primary key,
    type VARCHAR(255) not null,
    number VARCHAR(255) not null unique,
    employeeId int not null,
    foreign key (employeeId) references employees (id) on delete cascade
);

