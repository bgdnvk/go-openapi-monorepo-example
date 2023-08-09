https://github.com/deepmap/oapi-codegen 

# gen example:  
oapi-codegen -generate "chipo-example\services\swag2\api\v1\openapi>     -server,types" -package petstore petstore-expanded.yaml > petstore.gen.go  

u can find more in api/v1/openapispec