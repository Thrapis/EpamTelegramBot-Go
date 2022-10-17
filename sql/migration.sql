CREATE TABLE `users`
(
    id bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (id)
);


INSERT INTO `users` (`id`, `name`) VALUES (1, 'usr0'), (2, 'usr1');