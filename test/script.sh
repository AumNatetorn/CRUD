#Create Database
CREATE DATABASE customer

#Create Tables

CREATE TABLE profile (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    age INT
);


# Curl api Create Customer
curl -X POST http://localhost:1323/customers \
-H "Content-Type: application/json" \
-d '{"name": "natetorntest", "age": 26}'

#Curl api Update Customer
curl -X PUT http://localhost:1323/customers \
-H "Content-Type: application/json" \
-d '{"id":12,"name": "natetorntest2", "age": 22}'

#Curl api Delete Customer
curl -X DELETE http://localhost:1323/customers/12 \
-H "Content-Type: application/json" \
-d '{}'

#Curl api Find Customer
curl -X GET http://localhost:1323/customers/13 \
-H "Content-Type: application/json" \
-d '{}'

