# Student Api
- Make a cassandra docker container by the following command
```
make create_db
```
- Wait for 2-3 minutes for the container to activate and then open the cassandra shell by typing following in the terminal
```
make db_shell
```
- Now inside the cassandra shell, make the keyspace and students table by following commands
```
CREATE KEYSPACE student_api
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE emps (
  id uuid PRIMARY KEY,
  name text,
  age int,
  class text,
  branch text
);
```
- Run the following commands in sequence to run the project 
```
make build
make test
make run
```
- Finally, you will find the documentation at 127.0.0.1:8080/docs
