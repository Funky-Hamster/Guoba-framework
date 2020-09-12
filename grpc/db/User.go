package db
/*
 * CREATE TABLE user_tb(id INT PRIMARY KEY AUTO_INCREMENT
  ,username VARCHAR(30) DEFAULT ""
  ,password VARCHAR(255) NOT NULL
  ,age int
  , sex varchar(2) default "男"
 );
*/
type User struct{
	Id int32 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Token string `db:"token" json:"token"`
}
