
1. mysql 초기화 
create database zoyi_test;
create user zoyi@localhost identified by 'zoyi11';
grant all privileges on zoyi_test.* to zoyi@localhost with grant option;


CREATE TABLE `tb_key` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `name_index` (`name`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2
;


CREATE TABLE `tb_translation` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`keyId` VARCHAR(50) NOT NULL,
	`locale` VARCHAR(50) NULL DEFAULT NULL,
	`value` TEXT NULL DEFAULT NULL,
		PRIMARY KEY (`id`)
)
;

