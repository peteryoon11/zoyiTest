
1. mysql 초기화 및 테이블 생성 
```
create database zoyi_test;
create user zoyi@localhost identified by 'zoyi11';
grant all privileges on zoyi_test.* to zoyi@localhost with grant option;


CREATE TABLE `tb_key` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `name` (`name`),
	INDEX `name_index` (`name`)
)


CREATE TABLE `tb_translation` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`keyId` VARCHAR(50) NOT NULL,
	`locale` VARCHAR(50) NULL DEFAULT NULL,
	`value` TEXT NULL DEFAULT NULL,
		PRIMARY KEY (`id`)
)
;
```

2. 아래와 같이 이동 후에 실행 하시 거나 바이너리 파일을 이용해서 사용 
```
cd src 

go run main.go 

or 

cd src 

go build 

chmod 755 main 


./main or ./main.exe

```

3. localhost:8090 
