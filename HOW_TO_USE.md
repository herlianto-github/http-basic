# How to use

- Create Database

```sql
mysql -u 'database name' -p 'database password'
create database 'database name'

CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(5) PRIMARY,
  `name` varchar(45) ,
  `email` varchar(45) UNIQUE,
  `password` varchar(45)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(1, 'User 1', 'user1@email.com', '1234');
```

- Create your .env file

```go
APP_PORT:8000
DB_DRIVER:"mysql"
DB_ADDRESS:"127.0.0.1"
DB_PORT:3306
DB_USERNAME:"database username"
DB_PASSWORD:"database password"
DB_NAME:"database name"
```

- go run main.go

```sh
server started at localhost:8000
```