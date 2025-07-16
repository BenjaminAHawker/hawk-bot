CREATE TABLE IF NOT EXISTS requests (
	id SERIAL PRIMARY KEY,
	request_type_id INTEGER NOT NULL REFERENCES request_types(id),
	user_requested INTEGER NOT NULL REFERENCES users(id),
	status_id INTEGER NOT NULL DEFAULT 0 REFERENCES request_status_types(id),
	processed_by INTEGER REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_requests_user_requested ON requests(user_requested);