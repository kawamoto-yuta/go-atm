CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `name` varchar(25) NOT NULL COMMENT 'ユーザー名',
  `account_num` varchar(7) UNIQUE NOT NULL COMMENT '口座番号',
  `password` varchar(100) NOT NULL COMMENT 'パスワード',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime,
  PRIMARY KEY (`id`),
  INDEX id_index(id)
) DEFAULT CHARSET=utf8mb4;
