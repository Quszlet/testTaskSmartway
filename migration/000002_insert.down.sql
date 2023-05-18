DELETE FROM employees;
ALTER SEQUENCE employees_id_seq RESTART WITH 1;

DELETE FROM passports;
ALTER SEQUENCE passports_id_seq RESTART WITH 1;

DELETE FROM departments;
ALTER SEQUENCE departments_id_seq RESTART WITH 1;

DELETE FROM companies;
ALTER SEQUENCE companies_id_seq RESTART WITH 1;