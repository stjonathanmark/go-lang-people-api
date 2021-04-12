# Golang People API
Golang demo API that does CRUD operations for a simple **Person** model using MS SQL Server as the database. 

### Third Party Modules
1. [github.com/gorilla/mux](https://github.com/gorilla/mux) v1.8.0
2. [gorm.io/driver/sqlserver](https://github.com/go-gorm/sqlserver) v1.0.7
3. [gorm.io/gorm](https://github.com/go-gorm) v1.21.6

### Person Model
```golang
type Person struct {
  Id         int64 
  Deleted    bool  
  FirstName  string
  MiddleName string
  LastName   string
}
```
## How to run using Docker containers
1. Download and install [**Docker Desktop**](https://www.docker.com/products/docker-desktop) for your operating system
2. If neccessary **Switch to Linux Containers** in **Docker Desktop**
3. Open command line to of your choice to root directory for **Golang People API** app where ever it is located on your computer
4. Run the command **docker-compose up**
    * Due to **MS SQL Server** taking a while to start up, the **Golang People API** will likely have to attempt to connect multiple times before doing so successfully
    * By default the **Golang People API** makes 3 attempts 60 seconds apart to connect to **MS SQL Server**
    * You can modify the amount of attempts made by editing the **"MAX_DB_CONN_ATTEMPTS"** environmental variable in the [**"docker-compose.yml"**](https://github.com/stjonathanmark/go-lang-people-api/blob/master/docker-compose.yml) file
    * You can modify the amount of seconds between each attempt by editing the **"SECONDS_BTW_DB_CONN_ATTEMPTS"** environmental variable in the [**"docker-compose.yml"**](https://github.com/stjonathanmark/go-lang-people-api/blob/master/docker-compose.yml) file
5. Download and install [**Postman**](https://www.postman.com/downloads/) or any other API client of your choice
6. Run the following routes:
    * GET http://localhost:3000/api/person to get all people
    * GET http://localhost:3000/api/person/{id} to get a single person with specified id
    * POST http://localhost:3000/api/person to create a person
    ```json
    {
      "deleted": false,
      "firstName": "John",
      "middleName": null,
      "lastName": "Doe"
    }
    ```
    * PUT http://localhost:3000/api/person/{id} to update a person with specified id
    ```json
    {
      "deleted": false,
      "firstName": "John",
      "middleName": "William",
      "lastName": "Doe"
    }
    ```
    * DELETE http://localhost:3000/api/person/{id} to delete a person with specified id
7. Enjoy!!!!!!!
