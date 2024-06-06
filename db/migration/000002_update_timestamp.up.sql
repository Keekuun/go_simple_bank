ALTER TABLE accounts
    ALTER COLUMN created_at SET DATA TYPE timestamptz;

ALTER TABLE entries
    ALTER COLUMN created_at SET DATA TYPE timestamptz;

ALTER TABLE transfers
    ALTER COLUMN created_at SET DATA TYPE timestamptz;