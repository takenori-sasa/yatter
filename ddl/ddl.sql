CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint NOT NULL,
  `content` text NOT NULL,
  `url` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX idx_account_id (account_id),
  CONSTRAINT fk_status_account_id FOREIGN KEY (account_id) REFERENCES  account (id) ON DELETE CASCADE
);
CREATE TABLE `relationship` (
  `following_id` bigint(20) NOT NULL,
  `followed_id` bigint(20) NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`following_id`, `followed_id`),
  CONSTRAINT `fk_relationship_following_id` FOREIGN KEY (`following_id`) REFERENCES `account` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_relationship_followed_id` FOREIGN KEY (`followed_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
);
