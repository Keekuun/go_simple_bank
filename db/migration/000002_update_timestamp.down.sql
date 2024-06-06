ALTER TABLE accounts
    ALTER COLUMN created_at SET DATA TYPE timestamp;

ALTER TABLE entries
    ALTER COLUMN created_at SET DATA TYPE timestamp;

ALTER TABLE transfers
    ALTER COLUMN created_at SET DATA TYPE timestamp;