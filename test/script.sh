#Create Database
CREATE DATABASE customer

#Create Tables

CREATE TABLE profile (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT,
    name VARCHAR(255),
    age INT
);


# Curl api Create Customer
curl -X POST http://localhost:1323/customer \
-H "Content-Type: application/json" \
-d '{"accountID": 123, "name": "natetorntest", "age": 26}'

#Curl api Update Customer
curl -X PUT http://localhost:1323/customer \
-H "Content-Type: application/json" \
-d '{"id":10,"accountID": 123, "name": "natetorntest", "age": 26}'

#Curl api Delete Customer
curl -X DELETE http://localhost:1323/customer/10 \
-H "Content-Type: application/json" \
-d '{}'

#Curl api Find Customer
curl -X GET http://localhost:1323/customer/11 \
-H "Content-Type: application/json" \
-d '{}'

