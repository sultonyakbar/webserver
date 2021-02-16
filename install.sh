#!/bin/bash
apt-get update
apt-get install build-essential software-properties-common -y
#Instalasi Apache , PHP , Mariadb
apt install apache2 -y
apt install mariadb-server -y
sudo mysql -u root -p
CREATE DATABASE wpdatabase;
CREATE USER 'wpuser'@'localhost' IDENTIFIED BY 'new_password_here';
GRANT ALL PRIVILEGES ON wpdatabase . * TO 'wpuser'@'localhost';
FLUSH PRIVILEGES;
EXIT;

apt install php libapache2-mod-php php-mysql php-curl php-gd php-mbstring php-xml php-xmlrpc php-soap php-intl php-zip -y

#Instalasi Wordpress
cd /tmp
wget https://wordpress.org/latest.tar.gz
tar -xvzf latest.tar.gz
sudo mv wordpress /var/www/html/sultonyakbar.my.id
sudo chown -R www-data:www-data /var/www/html/sultonyakbar.my.id/
sudo chmod -R 755 /var/www/html/sultonyakbar.my.id/
cp apps.conf /etc/apache2/sites-available/
a2ensite apps.conf
a2dissite 000-default.conf
sudo a2enmod rewrite
sudo a2enmod dir
sudo a2enmod ssl
sudo a2enmod headers
sudo a2enmod http2
systemctl restart apache2

#Instalasi SSL Letsencrypt
add-apt-repository ppa:certbot/certbot -y
apt install python-certbot-apache -y 
certbot --apache -d sultonyakbar.my.id -d www.sultonyakbar.my.id
