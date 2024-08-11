create or replace package test2_api is
	procedure get (
		p_id in test2.id%type 
		,p_row out test2%rowtype
	);

	procedure insert (
		p_id in test2.id%type 
		,p_name in test2.name%type 
		,p_age in test2.age%type 
		,p_row out test2%rowtype
	);

	procedure update (
		p_id in test2.id%type 
		,p_name in test2.name%type 
		,p_age in test2.age%type 
		,p_row out test2%rowtype
	);

end;

create or replace package body test2_api is
	procedure get (
		p_id in test2.id%type 
		,p_row out test2%rowtype
	) is
		lrow test2%rowtype;
	begin
		select * into lrow from test2 where id = p_id;
		p_row := lrow;
	end get;

	procedure insert (
		p_id in test2.id%type 
		,p_name in test2.name%type 
		,p_age in test2.age%type 
		,p_row out test2%rowtype
	) is
		lrow test2%rowtype;
	begin
		insert into test2 (id, name, age) values (p_id, p_name, p_age)
		returning * into lrow;
		p_row := lrow;
	end insert;

	procedure update (
		p_id in test2.id%type 
		,p_name in test2.name%type 
		,p_age in test2.age%type 
		,p_row out test2%rowtype
	) is
		lrow test2%rowtype;
	begin
		update test2 set name = p_name, age = p_age where id = p_id	returning * into lrow;
		p_row := lrow;
	end update;

end;