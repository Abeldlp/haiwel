-- Drop the foreign key constraint
ALTER TABLE users
    DROP CONSTRAINT fk_user_ranking_id;

-- Rename the column back to its original name
ALTER TABLE users
    RENAME COLUMN user_raking_id TO rank;

-- Drop the table user_rankings
DROP TABLE user_rankings;