-- global state should only be updated by main character
-- since this intended to be a single player game, player is included in global state
CREATE TABLE IF NOT EXISTS global (
    -- game state
    paused          BOOL NOT NULL,
    should_close    BOOL NOT NULL,


    -- player state
    max_health      INTEGER NOT NULL,                   -- max health of player, can be changed
    curr_health     INTEGER NOT NULL,                   -- current health
    pos_x           INTEGER NOT NULL,
    pos_y           INTEGER NOT NULL,
    direction       INTEGER NOT NULL,                   -- direction, should be in degree
    facing          INTEGER NOT NULL                    -- direction, but where the player is facing
);


-- object should be bullet or enemy
-- state of object should provide just enough information for drawing it on screen
CREATE TABLE IF NOT EXISTS object (
    -- id, other fields that shouldn't be changed should be passed as process argument
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    -- state related to drawing but will not change
    max_health      INTEGER NOT NULL,                   -- max health
    show_health     BOOL NOT NULL,                      -- if show health

    -- state of the object not related to drawing
    should_close    BOOL NOT NULL DEFAULT FALSE,        -- should close

    -- current state
    curr_health     INTEGER NOT NULL,                   -- current health
    pos_x           INTEGER NOT NULL,
    pos_y           INTEGER NOT NULL,
    direction       INTEGER NOT NULL,                   -- direction, should be in degree
    facing          INTEGER NOT NULL                    -- direction, but where the object is facing
);


