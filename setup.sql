ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
CREATE DATABASE IF NOT EXISTS puppets;
use puppets;
CREATE TABLE IF NOT EXISTS set1(
	name VARCHAR(120),
	os VARCHAR(120),
	ip VARCHAR(120)
);
