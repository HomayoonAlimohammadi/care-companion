-- name: GetCareSeekerByName :one
select * from care_seeker
where name = $1; 

-- name: GetCareSeekerByID :one
select * from care_seeker
where id = $1; 

-- name: UpdateCareSeeker :exec
update care_seeker 
set name = $1 where id = $2;

-- name: DeleteCareSeekerByID :exec
delete from care_seeker where id = $1;

-- name: DeleteCareSeekerByName :exec
delete from care_seeker where name = $1;

-- name: CreateCareSeeker :exec
insert into care_seeker (name)
values ($1);;

