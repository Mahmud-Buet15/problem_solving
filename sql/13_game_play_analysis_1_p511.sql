-- PosegreSQL Solution
select 
    player_id
    , min(event_date) as first_login
from activity
group by 1


-- MySQL Solution
select 
    player_id
    , min(event_date) as first_login
from activity
group by 1
