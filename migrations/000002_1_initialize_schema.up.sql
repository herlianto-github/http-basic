CREATE TABLE IF NOT EXISTS `httpBasic`.`users` (
  `id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NULL,
  `email` VARCHAR(50) NULL,
  `password` VARCHAR(255) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);
