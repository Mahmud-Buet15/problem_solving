-- PosegreSQL Solution

with insurance_v2 as (
select 
    *
    , lat||'-'||lon as lat_lon
from insurance
)

, duplicate_investments as (
select 
    tiv_2015
    , count(1) as pid_count
from insurance_v2
group by 1
having count(1)>1
)

, duplicate_location as (
select 
    lat_lon
    , count(1) as pid_count
from insurance_v2
group by 1
having count(1)>1
)

select 
    round(sum(tbl1.tiv_2016)::numeric,2) as tiv_2016
from insurance_v2 tbl1 
left join duplicate_investments tbl2 on tbl1.tiv_2015=tbl2.tiv_2015
left join duplicate_location tbl3 on tbl1.lat_lon=tbl3.lat_lon
where 1=1
and tbl2.tiv_2015 is not null   --Excluding pids with unique investments
and tbl3.lat_lon is null        --filtering pids with non unique location


--===================================================
-- MySQL Solution
with insurance_v2 as (
select 
    *
    , concat(cast(lat as CHAR),'-',cast(lon as CHAR)) as lat_lon
from insurance
)

-- select * from insurance_v2

, duplicate_investments as (
select 
    tiv_2015
    , count(1) as pid_count
from insurance_v2
group by 1
having count(1)>1
)

, duplicate_location as (
select 
    lat_lon
    , count(1) as pid_count
from insurance_v2
group by 1
having count(1)>1
)

select 
    round(sum(tbl1.tiv_2016),2) as tiv_2016
from insurance_v2 tbl1 
left join duplicate_investments tbl2 on tbl1.tiv_2015=tbl2.tiv_2015
left join duplicate_location tbl3 on tbl1.lat_lon=tbl3.lat_lon
where 1=1
and tbl2.tiv_2015 is not null   
and tbl3.lat_lon is null     
;