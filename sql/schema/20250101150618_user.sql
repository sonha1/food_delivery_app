-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    username    VARCHAR(255) UNIQUE NOT NULL,
    name         VARCHAR(512)        NOT NULL,
    password     VARCHAR(512)        NOT NULL,
    email        VARCHAR(512) unique NOT NULL
) engine = InnoDB
  DEFAULT CHARSET = utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user`;
-- +goose StatementEnd
