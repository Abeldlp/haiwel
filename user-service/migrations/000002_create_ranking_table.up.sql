-- Create user_rankings table
CREATE TABLE user_rankings (
    id SERIAL PRIMARY KEY,
    game_id INTEGER
);

-- Create mock ranking for migration
INSERT INTO user_rankings (game_id) VALUES (1);

-- Update all the current users ranks to 1
UPDATE users SET rank=1;

-- Convert rank from an integer to an id on user_rankings table
ALTER TABLE users
RENAME COLUMN rank TO user_raking_id;

-- Add foreign key
ALTER TABLE users
ADD CONSTRAINT fk_user_ranking_id
FOREIGN KEY (user_raking_id)
REFERENCES user_rankings (id);

