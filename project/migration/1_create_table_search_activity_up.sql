CREATE TABLE IF NOT EXISTS search_activity (
	id            BIGSERIAL PRIMARY KEY,
	activity      VARCHAR(255) NOT NULL,
	success       BOOLEAN NOT NULL,
	message 	  VARCHAR(255) NOT NULL,
	json_data 	  TEXT,
	request_dt    TIMESTAMP
  );