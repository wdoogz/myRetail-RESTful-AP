### Useage

#### Step 1 (start mongodb):

```docker run -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -p 27017:27017 -v <path to repo>/mongodb:/data/db mongo```

#### Step 2 (start restful app)

``` go run main.go ```


The data looks like this in the mongodb instance

```
{ "_id" : ObjectId("6257b44d8fd4f5d2e3272f76"), "id" : "12954218", "current_price" : { "value" : 18.36, "currency_code" : "USD" } }
{ "_id" : ObjectId("6257b46e8fd4f5d2e3272f77"), "id" : "13264003", "current_price" : { "value" : 3.99, "currency_code" : "USD" } }
{ "_id" : ObjectId("6257b4848fd4f5d2e3272f78"), "id" : "54456119", "current_price" : { "value" : 15.55, "currency_code" : "USD" } }
{ "_id" : ObjectId("6257b4968fd4f5d2e3272f79"), "id" : "13860428", "current_price" : { "value" : 45, "currency_code" : "USD" } }
```
