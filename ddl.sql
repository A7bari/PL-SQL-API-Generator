CREATE TABLE employees (
    employee_id NUMBER(10) PRIMARY KEY,
    first_name  VARCHAR2(50),
    last_name   VARCHAR2(50) NOT NULL,
    email       VARCHAR2(100) NOT NULL UNIQUE,
    phone_number VARCHAR2(20),
    hire_date   DATE DEFAULT SYSDATE,
    job_id      VARCHAR2(10) NOT NULL,
    salary      NUMBER(8, 2),
    manager_id  NUMBER(10),
    department_id NUMBER(10)
);

CREATE TABLE departments (
    department_id NUMBER(10) PRIMARY KEY,
    department_name VARCHAR2(50) NOT NULL,
    manager_id NUMBER(10)
);

CREATE TABLE jobs (
    job_id VARCHAR2(10) PRIMARY KEY,
    job_title VARCHAR2(50) NOT NULL,
    min_salary NUMBER(8, 2),
    max_salary NUMBER(8, 2)
);


