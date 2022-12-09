# Test Stori José Antonio Limonta

This is an application intended to resolve technical challenge provided by Stori
The implementation of the solution is made with Go and Postgres as a database 
For emails we use a sendgrid Account
The application is dockerized to run the service and database in a single image


## Usage
To run the application execute the following command
```shell
    docker compose up
```
This will initialize the database service and the web service on port 8080.
The web service contains a single endpoint which you can call and send the email to where the report should be sent.

```curl
POST http://localhost:8080/statement
Accept: application/json

{
  "email": "pepelimonta@gmail.com"
}

```

The report is generated based on the file Test.csv

In the database connecting via <br>
host : localhost:5432<br>
user : postgres<br>
password: root<br>

The results of the Report can be seen in 2 tables report and transaction 

## Contributing

José Antonio Limonta Peddie