-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `permission`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `resource`   VARCHAR(255),
    `resources`  JSON,
    `role`       VARCHAR(255) NOT NULL,
    `action`     ENUM ('*', 'read', 'create', 'update', 'delete')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `permission`;
-- +goose StatementEnd
