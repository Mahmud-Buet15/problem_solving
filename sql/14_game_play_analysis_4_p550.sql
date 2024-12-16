-- PosegreSQL Solution

with first_login_info as (
select 
    player_id
    , min(event_date) as first_login
from activity
group by 1
)

, next_date_info as (
select 
    tbl1.*
    , tbl2.first_login
    , tbl2.first_login+1 as next_date
from activity tbl1
inner join first_login_info tbl2 on tbl1.player_id=tbl2.player_id
)

, check_consecutive as (
select 
    *
    , case when next_date=event_date then 1 else 0 end as is_consecutive
from next_date_info
)

-- select * from check_consecutive

, filtered as (
select 
    count(distinct player_id) as total_players
    , count(distinct case when is_consecutive=1 then player_id end) as total_players_filtered
from check_consecutive
)

-- select * from filtered

select 
    round(cast(total_players_filtered as decimal)/total_players,2) as fraction   
from filtered
;

/*
Notes:
- Postgres automatically converts to date if dateformat is ok
- postgres give integer divison as default when one integer is divided by another
*/
--===================================================

-- MySQL Solution
with first_login_info as (
select 
    player_id
    , min(event_date) as first_login
from activity
group by 1
)

, next_date_info as (
select 
    tbl1.*
    , tbl2.first_login
    , tbl2.first_login+1 as next_date
from activity tbl1
inner join first_login_info tbl2 on tbl1.player_id=tbl2.player_id
)

, check_consecutive as (
select 
    *
    , case when next_date=event_date then 1 else 0 end as is_consecutive
from next_date_info
)

-- select * from check_consecutive

, filtered as (
select 
    count(distinct player_id) as total_players
    , count(distinct case when is_consecutive=1 then player_id end) as total_players_filtered
from check_consecutive
)

-- select * from filtered

select 
    round(total_players_filtered /total_players,2) as fraction   
from filtered

/*
Notes:
- Mysql automatically converts to date if dateformat is ok
- Need to use DATE_ADD() function to add 1 day 
- gives float divison as default when one integer is divided by another
*/