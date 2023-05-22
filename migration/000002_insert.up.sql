INSERT INTO companies (name) VALUES ('smartway');
INSERT INTO companies (name) VALUES ('neoflex');
INSERT INTO companies (name) VALUES ('infotecs');

INSERT INTO departments (name, phone, companyId) VALUES ('Отдел кадров', '+79634476612', 1);
INSERT INTO departments (name, phone, companyId) VALUES ('Бухгалтерия', '+79567233333', 1);
INSERT INTO departments (name, phone, companyId) VALUES ('Отдел продаж', '+796345845693', 3);
INSERT INTO departments (name, phone, companyId) VALUES ('Отдел производства', '+79563271290', 1);
INSERT INTO departments (name, phone, companyId) VALUES ('Отдел производства', '+79563777777', 2);
INSERT INTO departments (name, phone, companyId) VALUES ('Отдел производства', '+79789321431', 3);

INSERT INTO employees (name, surname, phone, companyId, departmentId) 
VALUES ('Александр', 'Валуеев', '+79105467189', 1, 1);

INSERT INTO employees (name, surname, phone, companyId, departmentId) 
VALUES ('Владимир', 'Панькин', '+79912781211', 1, 1);

INSERT INTO employees (name, surname, phone, companyId, departmentId) 
VALUES ('Кирилл', 'Кирюхин', '+7814567888', 1, 4);

INSERT INTO employees (name, surname, phone, companyId, departmentId) 
VALUES ('Максим', 'Петрушков', '+79563726612', 2, 5);

INSERT INTO employees (name, surname, phone, companyId, departmentId) 
VALUES ('Марат', 'Болеев', '+79604567844', 3, 3);

INSERT INTO passports (type, number, employeeId) VALUES ('RU', '5834 657890', 1);
INSERT INTO passports (type, number, employeeId) VALUES ('RU', '6578 863890', 2);
INSERT INTO passports (type, number, employeeId) VALUES ('RU', '4300 621531', 3);
INSERT INTO passports (type, number, employeeId) VALUES ('RU', '1095 893231', 4);
INSERT INTO passports (type, number, employeeId) VALUES ('RU', '1111 503231', 5);