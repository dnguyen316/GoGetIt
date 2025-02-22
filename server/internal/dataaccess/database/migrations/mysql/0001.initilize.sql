CREATE TABLE IF NOT EXISTS accounts (
  account_id BIGINT UNSIGNED PRIMARY KEY,
  accountname VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS account_password (
  of_account_id UNSIGNED BIGINT PRIMARY KEY,
  hashed VARCHAR(128) NOT NULL,
  FOREIGN KEY (of_account_id) REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS download_tasks (
  task_id BIGINT PRIMARY KEY,
  of_account_id BIGINT UNSIGNED,
  download_type VARCHAR(128),
  url TEXT NOT NULL,
  download_status SMALLINT NOT NULL,
  metadata TEXT NOT NULL,
  FOREIGN KEY (of_account_id) REFERENCES accounts(account_id)
);

