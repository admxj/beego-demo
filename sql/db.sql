CREATE TABLE `guess` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `option` varchar(255) DEFAULT NULL,
  `img` varchar(255) DEFAULT NULL,
  `answer` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4

INSERT INTO `nginx`.`guess`(`id`, `name`, `age`, `option`, `img`, `answer`) VALUES (1, 'admxj', 24, '{\n  \"A\": \"王者荣耀\",\n  \"B\": \"QQ飞车\"\n}', '/img/cut/wangzherongyao.jpeg', 'A');
INSERT INTO `nginx`.`guess`(`id`, `name`, `age`, `option`, `img`, `answer`) VALUES (2, 'adm', 20, '{\n  \"A\": \"Ipad\",\n  \"B\": \"Iphone\"\n}', '/img/cut/ipad.jpg', 'A');