CREATE TABLE IF NOT EXISTS global (
    -- game state
    paused          BOOL NOT NULL,
    should_close    BOOL NOT NULL,                      -- if the game should be closed

    -- screen size
    screen_width    INTEGER NOT NULL,                   -- screen width
    screen_height   INTEGER NOT NULL                    -- screen height
);


CREATE TABLE IF NOT EXISTS player (
    -- player state
    max_health      FLOAT NOT NULL,                     -- max health of player, can be changed
    curr_health     FLOAT NOT NULL,                     -- current health
    pos_x           FLOAT NOT NULL,
    pos_y           FLOAT NOT NULL,
    size_x          FLOAT NOT NULL,
    size_y          FLOAT NOT NULL,
    direction       FLOAT NOT NULL,                     -- direction, should be in degree
    facing          FLOAT NOT NULL,                     -- direction, but where the player is facing
    status          INTEGER NOT NULL                    -- idle or moving(enum)
)


CREATE TABLE IF NOT EXISTS object (
    -- id and type
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    type            INTEGER NOT NULL,                   -- type of the object

    -- state related to drawing but will not change
    max_health      FLOAT NOT NULL,                     -- max health
    show_health     BOOL NOT NULL,                      -- if show health

    -- current state
    curr_health     FLOAT NOT NULL,                     -- current health
    pos_x           FLOAT NOT NULL,
    pos_y           FLOAT NOT NULL,
    size_x          FLOAT NOT NULL,
    size_y          FLOAT NOT NULL,
    direction       FLOAT NOT NULL,                     -- direction, should be in degree
    facing          FLOAT NOT NULL,                     -- direction, but where the object is facing
    status          INTEGER NOT NULL                    -- idle or moving(enum)
);


