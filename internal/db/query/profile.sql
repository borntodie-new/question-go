-- name: CreateProfile :one
INSERT INTO profiles (user_id,
                      real_name,
                      quote,
                      address)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: RetrieveProfile :one
SELECT u.*,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM profiles AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id
WHERE p.user_id = $1 LIMIT 1;

-- name: UpdateProfile :one
UPDATE profiles
SET real_name = $2,
    quote     = $3,
    address   = $4
WHERE user_id = $1 RETURNING *;
