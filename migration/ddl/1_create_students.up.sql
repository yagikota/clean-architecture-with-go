-- Active: 1661332216709@@127.0.0.1@3306@sample

CREATE SCHEMA IF NOT EXISTS `sample` DEFAULT CHARACTER SET utf8mb4 ;

USE `sample` ;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS students (
        id INT(11) AUTO_INCREMENT COMMENT "student ID",
        name VARCHAR(20) NOT NULL COMMENT "name",
        birthday DATE NOT NULL COMMENT "birthday",
        class INT(11) NOT NULL COMMENT "class number",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "student creation date",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "student updation date",
        deleted_at DATETIME COMMENT "student logical deletion date",
        PRIMARY KEY (id)
    ) ENGINE = INNODB COMMENT = 'student';
