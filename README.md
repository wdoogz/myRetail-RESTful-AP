### Useage

#### Step 1 (start mongodb):

```docker run -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -p 27017:27017 mongo```

#### Step 2 (start restful app)

To initilize the database with mock db use the `LOADDB=True` environment var.

``` MONGOUSER="<mongo user above> MONGOPASS="<mongo pass above>" go run main.go ```

##### Get Request
`curl localhost:9999/products/<id>`

##### PUT Request
```curl -X PUT localhost:9999/products/<id> -d <json output from above with new Value for price```