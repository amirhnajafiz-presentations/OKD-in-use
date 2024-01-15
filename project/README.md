# Project (MySQL Resolver)

In this project we are going to build and deploy a ```MySQL``` database resolver on __Kubernetes__ cluster.
This app gets a database url and credentials to check if the database can be reached in the namespace.

## How to use?

The docker image of this app is ```image-registry.openshift-image-registry.svc:5000/snappline-app-staging/okd-in-use-project```.
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

## Tagging

In order to send this image to OKD image registry you need to login to the cluster with your
credentials.

```shell
docker login -u <okd4-user> -p <okd4-token> <registery>/<namespace>
```

In our case we have ```image-registry.apps.private.okd4.teh-1.snappcloud.io/snappline-app-staging```.

Then tag your image like this:

```shell
docker tag amirhossein21/okd-in-use-project:v0.1.0 image-registry.apps.private.okd4.teh-1.snappcloud.io/snappline-app-staging/okd-in-use-project:v0.1.0
```
