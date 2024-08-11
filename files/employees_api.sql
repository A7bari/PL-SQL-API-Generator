create or replace package employees_api is
	procedure get (
		p_employee_id in employees.employee_id%type 
		,p_row out employees%rowtype
	);

	procedure insert (
		p_employee_id in employees.employee_id%type 
		,p_first_name in employees.first_name%type 
		,p_last_name in employees.last_name%type 
		,p_email in employees.email%type 
		,p_phone_number in employees.phone_number%type 
		,p_hire_date in employees.hire_date%type 
		,p_job_id in employees.job_id%type 
		,p_salary in employees.salary%type 
		,p_manager_id in employees.manager_id%type 
		,p_department_id in employees.department_id%type 
		,p_row out employees%rowtype
	);

	procedure update (
		p_employee_id in employees.employee_id%type 
		,p_first_name in employees.first_name%type 
		,p_last_name in employees.last_name%type 
		,p_email in employees.email%type 
		,p_phone_number in employees.phone_number%type 
		,p_hire_date in employees.hire_date%type 
		,p_job_id in employees.job_id%type 
		,p_salary in employees.salary%type 
		,p_manager_id in employees.manager_id%type 
		,p_department_id in employees.department_id%type 
		,p_row out employees%rowtype
	);

end;

create or replace package body employees_api is
	procedure get (
		p_employee_id in employees.employee_id%type 
		,p_row out employees%rowtype
	) is
		lrow employees%rowtype;
	begin
		select * into lrow from employees where employee_id = p_employee_id;
		p_row := lrow;
	end get;

	procedure insert (
		p_employee_id in employees.employee_id%type 
		,p_first_name in employees.first_name%type 
		,p_last_name in employees.last_name%type 
		,p_email in employees.email%type 
		,p_phone_number in employees.phone_number%type 
		,p_hire_date in employees.hire_date%type 
		,p_job_id in employees.job_id%type 
		,p_salary in employees.salary%type 
		,p_manager_id in employees.manager_id%type 
		,p_department_id in employees.department_id%type 
		,p_row out employees%rowtype
	) is
		lrow employees%rowtype;
	begin
		insert into employees (employee_id, first_name, last_name, email, phone_number, hire_date, job_id, salary, manager_id, department_id) values (p_employee_id, p_first_name, p_last_name, p_email, p_phone_number, p_hire_date, p_job_id, p_salary, p_manager_id, p_department_id)
		returning * into lrow;
		p_row := lrow;
	end insert;

	procedure update (
		p_employee_id in employees.employee_id%type 
		,p_first_name in employees.first_name%type 
		,p_last_name in employees.last_name%type 
		,p_email in employees.email%type 
		,p_phone_number in employees.phone_number%type 
		,p_hire_date in employees.hire_date%type 
		,p_job_id in employees.job_id%type 
		,p_salary in employees.salary%type 
		,p_manager_id in employees.manager_id%type 
		,p_department_id in employees.department_id%type 
		,p_row out employees%rowtype
	) is
		lrow employees%rowtype;
	begin
		update employees set first_name = p_first_name, last_name = p_last_name, email = p_email, phone_number = p_phone_number, hire_date = p_hire_date, job_id = p_job_id, salary = p_salary, manager_id = p_manager_id, department_id = p_department_id where employee_id = p_employee_id	returning * into lrow;
		p_row := lrow;
	end update;

end;