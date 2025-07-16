INSERT INTO request_types (description) VALUES ('audiobook') ON CONFLICT DO NOTHING;
INSERT INTO request_types (description) VALUES ('tv') ON CONFLICT DO NOTHING;
INSERT INTO request_types (description) VALUES ('movie') ON CONFLICT DO NOTHING;
INSERT INTO request_types (description) VALUES ('anime') ON CONFLICT DO NOTHING;
INSERT INTO request_types (description) VALUES ('ebook') ON CONFLICT DO NOTHING;