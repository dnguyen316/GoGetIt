CREATE TABLE IF NOT EXISTS users (
  user_id BIGINT UNSIGNED PRIMARY KEY,
  username VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_password (
  of_user_id UNSIGNED BIGINT PRIMARY KEY,
  hashed VARCHAR(128) NOT NULL,
  FOREIGN KEY (of_user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS download_tasks (
  task_id BIGINT PRIMARY KEY,
  of_user_id BIGINT UNSIGNED,
  download_type VARCHAR(128),
  url TEXT NOT NULL,
  download_status SMALLINT NOT NULL,
  metadata TEXT NOT NULL,
  FOREIGN KEY (of_user_id) REFERENCES users(user_id)
);

