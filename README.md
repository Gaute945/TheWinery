# TheWinery

## TODO (in priority)
1. upload page
2. edit page
3. proper units on data
4. images
5. user accounts
6. real data
7. sorting

## setup (dev)

### mariadb install
after mariadb install run: sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql   
systemctl start mariadb

### db and user setup
sudo mariadb   
CREATE DATABASE TheWinery   
CREATE USER 'JohnDoe'@'localhost' IDENTIFIED BY 'some_pass';   
GRANT ALL PRIVILEGES ON mydb.* TO 'JohnDoe'@'localhost';

### source the example db
mariadb -u JohnDoe -p   
USE TheWinery;   
SOURCE /full/path/to/dump.sql;

### .env config
nano .env   
DBUSER=JohnDoe   
DBPASS=some_pass   
DBNAME=TheWinery

### running
go run .   
localhost:8080

### misc
'/c' to exit from '->'   
'exit' to exit normally