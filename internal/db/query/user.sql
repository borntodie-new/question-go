-- name: CreateUser :one
INSERT INTO users (username,
                   nickname,
                   hashed_password)
VALUES ($1, $2, $3) RETURNING *;


-- name: RetrieveUserByUsername :one
SELECT u.*,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id
WHERE u.username = $1 LIMIT 1;

-- name: RetrieveUserByID :one
SELECT u.*,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id
WHERE u.id = $1 LIMIT 1;

-- name: QueryUsers :many
SELECT u.*,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id LIMIT $1
OFFSET $2;

-- name: DeleteUser :one
UPDATE users
SET status = $2
WHERE id = $1 RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET hashed_password = $2
WHERE id = $1 RETURNING *;

-- name: UpdateAdmin :one
UPDATE users
SET is_super = $2
WHERE id = $1 RETURNING *;