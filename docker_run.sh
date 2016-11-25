#
docker run -ti --name cache-server-go -v $(pwd)/src:/go/src -p 3000:8080 goenvironment:1.0
