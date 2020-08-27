package db

/**
@func 对应数据库user_tb 表
@author:柠檬191104
@date:2020/04/26
*/

/*
 * CREATE TABLE user_tb(id INT PRIMARY KEY AUTO_INCREMENT
  ,username VARCHAR(30) DEFAULT ""
  ,password VARCHAR(255) NOT NULL
  ,age int
  , sex varchar(2) default "男"
 );
*/
type User struct{
	Id int `db:"id" json:"id"`
	name string `db:"name" json:"name"`
	token string `db:"token" json:"token"`
}
