# Project (MySQL Resolver)

In this project we are going to build and deploy a ```MySQL``` database resolver on __Kubernetes__ cluster.
This app gets a database url and credentials to check if the database can be reached in the namespace.

## How to use?

The docker image of this app is ```amirhossein21/okd-in-use-project```.
You can use the manifests in the deployment directory to set up this service on kubernetes cluster.

After setting up the service, send a post request to resolve your database:

```shell
curl -i -X POST \ 
  -H 'Content-type: application/json' \ 
  -d '{"host":"127.0.0.1", "port":3306}'
  http://serviceip/api/resolve
```

Or you can just open the default address on your browser to use the UI panel.

### ENVs

The only environment variable for this project is __HTTP_PORT__, which is the port of our service.
The value should be a number between 1 and 23456.
