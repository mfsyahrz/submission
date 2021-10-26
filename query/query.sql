select 
    u1.id as "ID", 
    u1.username as "UserName", 
    (select username from "user" u2 where u2.id = u1.parent) as "ParentUserName"
from "user" u1;