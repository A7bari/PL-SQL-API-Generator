create or replace package departments_api is
	procedure get (
		p_department_id in departments.department_id%type 
		,p_row out departments%rowtype
	);

	procedure insert (
		p_department_id in departments.department_id%type 
		,p_department_name in departments.department_name%type 
		,p_manager_id in departments.manager_id%type 
		,p_row out departments%rowtype
	);

	procedure update (
		p_department_id in departments.department_id%type 
		,p_department_name in departments.department_name%type 
		,p_manager_id in departments.manager_id%type 
		,p_row out departments%rowtype
	);

end;

create or replace package body departments_api is
	procedure get (
		p_department_id in departments.department_id%type 
		,p_row out departments%rowtype
	) is
		lrow departments%rowtype;
	begin
		select * into lrow from departments where department_id = p_department_id;
		p_row := lrow;
	end get;

	procedure insert (
		p_department_id in departments.department_id%type 
		,p_department_name in departments.department_name%type 
		,p_manager_id in departments.manager_id%type 
		,p_row out departments%rowtype
	) is
		lrow departments%rowtype;
	begin
		insert into departments (department_id, department_name, manager_id) values (p_department_id, p_department_name, p_manager_id)
		returning * into lrow;
		p_row := lrow;
	end insert;

	procedure update (
		p_department_id in departments.department_id%type 
		,p_department_name in departments.department_name%type 
		,p_manager_id in departments.manager_id%type 
		,p_row out departments%rowtype
	) is
		lrow departments%rowtype;
	begin
		update departments set department_name = p_department_name, manager_id = p_manager_id where department_id = p_department_id	returning * into lrow;
		p_row := lrow;
	end update;

end;