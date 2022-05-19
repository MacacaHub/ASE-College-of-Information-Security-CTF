USE server_db;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL PRIMARY KEY,
  `username` text NOT NULL,
  `password` text NOT NULL
);

CREATE TABLE IF NOT EXISTS `flag` (
  `id` int NOT NULL PRIMARY KEY,
  `content` text NOT NULL
);

INSERT IGNORE INTO `users` (`id`, `username`, `password`) VALUES 
(1, 'macaca', 'm4c4c4_CC'),
(2, 'admin', '4dmin_au4a83'),
(3, 'user', '_us3R_us3R_');

INSERT IGNORE INTO `flag` (`id`, `content`) VALUES (1, 'FLAG{sql_1nject10n_by_un10n_select}');
