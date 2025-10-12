CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `role` varchar(255),
  `booster` bool,
  `created_at` timestamp
);

CREATE TABLE `role` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `color` varchar(255),
  `created_at` timestamp
);

ALTER TABLE `role` ADD FOREIGN KEY (`name`) REFERENCES `users` (`role`);
