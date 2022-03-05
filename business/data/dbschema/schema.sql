-- Description: Create table daily_combination
CREATE TABLE IF NOT EXISTS daily_combination (
	id SERIAL,
	combination TEXT,
	date_created TIMESTAMP,
	date_updated TIMESTAMP
);