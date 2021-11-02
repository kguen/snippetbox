--
-- Table structure for table `snippets`
--

CREATE TABLE `snippets` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created` datetime NOT NULL,
  `expires` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `snippets`
  ADD KEY `idx_snippets_created` (`created`) USING BTREE;

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `hashed_password` char(97) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `users`
  ADD UNIQUE KEY `users_uc_email` (`email`);

INSERT INTO `users` (`id`, `name`, `email`, `hashed_password`, `created`) VALUES(
    1,
    'Khoa Nguyễn',
    'nanhkhoa460@gmail.com',
    '$argon2id$v=19$m=65536,t=1,p=2$et0ceTGOYmW1CU64gREVNQ$RKhbtlYC1NiIxpvwvYabKapq2tkNL2jBVa/VqlSx6+Q',
    '2021-11-02 14:56:21');
