#!/bin/bash
if (whoami != root)
  then echo "Please run as root"

  else (setup)
fi

setup() { 
    sudo apt install mysql-server -y
    sudo apt install golang
    sudo go get github.com/go-sql-driver/mysql
    sudo go get github.com/olekukonko/tablewriter
    sudo go get github.com/CrowdSurge/banner
    sudo mysql -u root -p < setup.sql
}

