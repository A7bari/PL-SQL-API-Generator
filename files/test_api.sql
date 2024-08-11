create or replace package test_api is
	procedure get (
		p_id in test.id%type 
		,p_row out test%rowtype
	);

	procedure insert (
		p_id in test.id%type 
		,p_name in test.name%type 
		,p_age in test.age%type 
		,p_row out test%rowtype
	);

	procedure update (
		p_id in test.id%type 
		,p_name in test.name%type 
		,p_age in test.age%type 
		,p_row out test%rowtype
	);

end;

create or replace package body test_api is
	procedure get (
		p_id in test.id%type 
		,p_row out test%rowtype
	) is
		lrow test%rowtype;
	begin
		select * into lrow from test where id = p_id;
		p_row := lrow;
	end get;

	procedure insert (
		p_id in test.id%type 
		,p_name in test.name%type 
		,p_age in test.age%type 
		,p_row out test%rowtype
	) is
		lrow test%rowtype;
	begin
		insert into test (id, name, age) values (p_id, p_name, p_age)
		returning * into lrow;
		p_row := lrow;
	end insert;

	procedure update (
		p_id in test.id%type 
		,p_name in test.name%type 
		,p_age in test.age%type 
		,p_row out test%rowtype
	) is
		lrow test%rowtype;
	begin
		update test set name = p_name, age = p_age where id = p_id	returning * into lrow;
		p_row := lrow;
	end update;

end;