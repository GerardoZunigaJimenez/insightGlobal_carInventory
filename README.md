# insightGlobal_carInventory

## Steps to run it locally

Install dependencies 
```    
go mod tidy
```

Install mockery with brew
```
brew install mockery
```

Run Mockery to update mocks
```
mockery
```

Install Postgres & Run this dumpy.sql file to create the database
```

```


## REST Operations
### **Health check Endpoint**
    ``` 
    http://localhost:8080/healthcheck
    ``` 


### **CARS Endpoint Rest Calls**
- Create a new car
     ``` 
  POST http://localhost:8080/cars
  
  curl --location --request POST 'http://localhost:8080/cars' \
    --header 'Content-Type: text/plain' \
    --data-raw '{
    "make": "Honda",
    "model": "Civic",
    "year": 2023,
    "color": "Blue",
    "vin": "2HGFE2F59PH123457",
    "mileage": 8000,
    "price": 21950.0,
    "disabled": false
    },' 
    ```


- Retrieve a car by ID
  ``` 
  GET http://localhost:8080/:id
  
  http://localhost:8080/cars/550e8400-e29b-41d4-a716-446655440000 
  ```


- List all cars by PageSize and PageNumber
  ``` 
  GET http://localhost:8080/cars?pageSize=:pageSize&pageNumber=:pageNumber 

  http://localhost:8080/cars?pageSize=1&pageNumber=2
  ```
- Update an existing car
    ``` 
  UPDATE http://localhost:8080/cars/:id

  curl --location --request PUT 'http://localhost:8080/cars/04c509aa-7038-4b6b-bf51-c6167118188a' \
    --header 'Content-Type: text/plain' \
    --data-raw '{
    "id": "04c509aa-7038-4b6b-bf51-c6167118188a",
    "make": "Honda",
    "model": "Civic",
    "year": 2023,
    "color": "Blue and Red",
    "vin": "2HGFE2F59PH123457",
    "mileage": 8000,
    "price": 21950.0,
    "disabled": false
    },'
    ```
    
- Delete a car by ID
    ``` 
  Delete http://localhost:8080/cars/:id 
  
  curl --location --request DELETE 'http://localhost:8080/cars/550e8400-e29b-41d4-a716-446655440000'
  ```


