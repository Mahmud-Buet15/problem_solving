-- PosegreSQL Solution

with filtered_data as (
select 
    managerid
    , count(1) as total_reports
from employee
group by 1
having count(1)>=5   --getting manager with >=5 reports
)

select 
    tbl1.name
from employee tbl1
inner join filtered_data tbl2 on tbl1.id=tbl2.managerid


--===================================================
-- MySQL Solution
with filtered_data as (
select 
    managerid
    , count(1) as total_reports
from employee
group by 1
having count(1)>=5   
)

select 
    tbl1.name
from employee tbl1
inner join filtered_data tbl2 on tbl1.id=tbl2.managerid


