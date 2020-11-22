#!/bin/bash
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

sudo apt install mysql-server -y
sudo apt install golang
sudo go get github.com/go-sql-driver/mysql
sudo go get github.com/olekukonko/tablewriter
sudo go get github.com/CrowdSurge/banner
sudo mysql -u root -p < setup.sql


