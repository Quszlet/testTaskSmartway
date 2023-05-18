INSERT INTO passports (type, number) VALUES ('RU', '5834 657890');
INSERT INTO passports (type, number) VALUES ('RU', '6578 863890');
INSERT INTO passports (type, number) VALUES ('RU', '4300 621531');
INSERT INTO passports (type, number) VALUES ('RU', '1095 893231');
INSERT INTO passports (type, number) VALUES ('RU', '1111 503231');

INSERT INTO departments (name, phone) VALUES ('Отдел кадров', '+79634476612');
INSERT INTO departments (name, phone) VALUES ('Бухгалтерия', '+79567233333');
INSERT INTO departments (name, phone) VALUES ('Отдел продаж', '+796345845693');

INSERT INTO companies (name) VALUES ('Smartway');
INSERT INTO companies (name) VALUES ('Neoflex');
INSERT INTO companies (name) VALUES ('Infotecs');

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
VALUES ('Александр', 'Валуеев', '+79105467189', 1, 1, 1);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
VALUES ('Владимир', 'Панькин', '+79912781211', 1, 2, 2);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
VALUES ('Кирилл', 'Кирюхин', '+7814567888', 1, 3, 3);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
VALUES ('Максим', 'Петрушков', '+79563726612', 1, 4, 1);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
VALUES ('Марат', 'Болеев', '+79604567844', 2, 5, 2);