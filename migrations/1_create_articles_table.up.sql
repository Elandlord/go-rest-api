CREATE TABLE `articles` (
  `Id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Content` longtext,
  `Description` longtext,
  `Title` longtext,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Id` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;