#!/bin/bash
apt-get update
apt-get install build-essential software-properties-common -y
#Instalasi Apache , PHP , Mariadb
apt install nginx -y
apt install mariadb-server -y
sudo mysql -u root -p
CREATE DATABASE wpdatabase;
CREATE USER 'wpuser'@'localhost' IDENTIFIED BY 'new_password_here';
GRANT ALL PRIVILEGES ON wpdatabase . * TO 'wpuser'@'localhost';
FLUSH PRIVILEGES;
EXIT;

apt install  php-mysql php-curl php-gd php-mbstring php-xml php-xmlrpc php-soap php-intl php-zip php-fpm -y
