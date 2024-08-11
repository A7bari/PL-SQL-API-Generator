create or replace package jobs_api is
	procedure get (
		p_job_id in jobs.job_id%type 
		,p_row out jobs%rowtype
	);

	procedure insert (
		p_job_id in jobs.job_id%type 
		,p_job_title in jobs.job_title%type 
		,p_min_salary in jobs.min_salary%type 
		,p_max_salary in jobs.max_salary%type 
		,p_row out jobs%rowtype
	);

	procedure update (
		p_job_id in jobs.job_id%type 
		,p_job_title in jobs.job_title%type 
		,p_min_salary in jobs.min_salary%type 
		,p_max_salary in jobs.max_salary%type 
		,p_row out jobs%rowtype
	);

end;

create or replace package body jobs_api is
	procedure get (
		p_job_id in jobs.job_id%type 
		,p_row out jobs%rowtype
	) is
		lrow jobs%rowtype;
	begin
		select * into lrow from jobs where job_id = p_job_id;
		p_row := lrow;
	end get;

	procedure insert (
		p_job_id in jobs.job_id%type 
		,p_job_title in jobs.job_title%type 
		,p_min_salary in jobs.min_salary%type 
		,p_max_salary in jobs.max_salary%type 
		,p_row out jobs%rowtype
	) is
		lrow jobs%rowtype;
	begin
		insert into jobs (job_id, job_title, min_salary, max_salary) values (p_job_id, p_job_title, p_min_salary, p_max_salary)
		returning * into lrow;
		p_row := lrow;
	end insert;

	procedure update (
		p_job_id in jobs.job_id%type 
		,p_job_title in jobs.job_title%type 
		,p_min_salary in jobs.min_salary%type 
		,p_max_salary in jobs.max_salary%type 
		,p_row out jobs%rowtype
	) is
		lrow jobs%rowtype;
	begin
		update jobs set job_title = p_job_title, min_salary = p_min_salary, max_salary = p_max_salary where job_id = p_job_id	returning * into lrow;
		p_row := lrow;
	end update;

end;