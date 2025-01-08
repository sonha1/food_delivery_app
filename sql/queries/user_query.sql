-- name: GetAllUser :many
SELECT *
FROM `user`
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: FindUserByUsername :one
SELECT *
FROM `user`
where username = ?;

-- name: CreateUser :execresult
INSERT INTO `user` (username, password, name, email)
VALUES (?, ?, ? ,?) ;

-- name: FindUserById :one
SELECT *
FROM `user`
where id = ?;
