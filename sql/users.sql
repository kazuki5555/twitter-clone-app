CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `screen_name` varchar(255) NOT NULL DEFAULT '' COMMENT '表示名',
  `screen_id` varchar(15) NOT NULL  DEFAULT '' COMMENT '表示id',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'email',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'パスワード',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `screen_id` (`screen_id`),
  UNIQUE KEY `screen_id_2` (`screen_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;