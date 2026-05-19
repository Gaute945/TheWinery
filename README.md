# TheWinery

## TODO (in priority)

1. edit changes nil fields to random memory or "nil"
2. non orderd ids  
3. Fining functions only prints for now  
4. user accounts / auth
5. real data
6. sorting

## setup (dev)

### mariadb install

after mariadb install run: sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql  
systemctl start mariadb  

### db and user setup

sudo mariadb  
CREATE DATABASE TheWinery;  
CREATE USER 'JohnDoe'@'localhost' IDENTIFIED BY 'some_pass';  
GRANT ALL PRIVILEGES ON TheWinery.* TO 'JohnDoe'@'localhost';  

### source the example db

mariadb -u JohnDoe -p  
USE TheWinery;  
SOURCE create-tables.sql;  

### .env config

nano .env  
DBUSER=JohnDoe  
DBPASS=some_pass  

### running

go run .  
localhost:8080  

### info

I would appreciate having a dialog if you want to make a pull request, Discord: gauteg  
'/c' to exit from '->'  
'exit' to exit normally  
