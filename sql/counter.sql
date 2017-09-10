CREATE DATABASE if NOT EXISTS Ngdule ;
USE Ngdule;
# 首先要把用户模型设计好
CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(30) NOT NULL ,
  password VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS counter_competitor (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS counter_judger (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS counter_vote (
  id INT PRIMARY KEY AUTO_INCREMENT,
  uid INT NOT NULL ,
  cid INT NOT NULL ,
  score int NOT NULL ,
  type VARCHAR(30) NOT NULL
);
