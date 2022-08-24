-- Active: 1661332216709@@127.0.0.1@3306@sample

CREATE SCHEMA IF NOT EXISTS `sample` DEFAULT CHARACTER SET utf8mb4 ;

USE `sample` ;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS students (
        id INT(11) AUTO_INCREMENT COMMENT "学籍番号",
        name VARCHAR(8) NOT NULL COMMENT "氏名",
        birthday DATE NOT NULL COMMENT "誕生日",
        class INT(11) NOT NULL COMMENT "クラス番号",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "ユーザー作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "ユーザー更新日時",
        deteled_at DATETIME COMMENT "ユーザー論理削除日時",
        PRIMARY KEY (id)
    ) ENGINE = INNODB COMMENT = '学生';