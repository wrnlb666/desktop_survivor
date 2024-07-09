-- name: NewGlobal :exec
INSERT INTO global
    (paused, should_close, screen_width, screen_height)
VALUES
    (FALSE, FALSE, ?, ?);


-- name: GetGlobal :one
SELECT *
FROM global
LIMIT 1;


-- name: SetGlobal :exec
UPDATE global
SET
    paused = ?,
    should_close = ?,
    screen_width = ?,
    screen_height = ?;


-- name: NewObject :one
INSERT INTO object (
    type,
    max_health,
    show_health,
    curr_health,
    pos_x,
    pos_y,
    size_x,
    size_y,
    direction,
    facing,
    status
)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id;


-- name: GetObject :one
SELECT *
FROM object
WHERE id = ?;


-- name: SetObject :exec
UPDATE object
SET
    max_health = ?,
    show_health = ?,
    curr_health = ?,
    pos_x = ?,
    pos_y = ?,
    size_x = ?,
    size_y = ?,
    direction = ?,
    facing = ?,
    status = ?
WHERE
    id = ?;


-- name: GetPlayer :one
SELECT *
FROM player
LIMIT 1;


-- name: SetPlayer :exec
UPDATE player
SET
    max_health = ?,
    curr_health = ?,
    pos_x = ?,
    pos_y = ?,
    size_x = ?,
    size_y = ?,
    direction = ?,
    facing = ?,
    status = ?;
