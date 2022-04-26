CREATE DATABASE IF NOT EXISTS development;
USE development;

CREATE TABLE IF NOT EXISTS guests
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `first_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    `last_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    `first_name_kana` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL,
    `last_name_kana` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL,
    `gender` int(11) DEFAULT '1',
    `country_code` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
    `tel` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
    `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    `birthday` date DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
    );

LOCK TABLES `guests` WRITE;
INSERT INTO guests VALUES (1,'First Name','Last Name','Fist Name Kana','Last Name Kana',1,'81','0312341234','guest@dev.local','2001-03-04', '2022-03-04 15:58:22','2022-03-10 12:05:48');
UNLOCK TABLES;