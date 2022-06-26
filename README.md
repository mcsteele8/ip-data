# ip-data
### Building & Running:
 - Configuration
   - `$export DOMAIN_API_KEY="your-api-key-to-api.ip2whois.com"` This is used to query data for domain names. [Docs](https://www.ip2whois.com/developers-api)
   - `$export IP_DATA_SERVER_PORT=":80"` Default server port is 8081
 - Using Docker: 
   - [Install docker](https://docs.docker.com/get-docker/)
   - Build image `` docker build --tag docker-ip-data .``
   - Run image ``docker run docker-ip-data``
   - Url path: `localhost:8081/v1/data/{ip-or-domain}`
 - Using Go: 
   - [Install Go](https://golang.org/doc/install)
   - Terminal command: ``go run main.go``
   - Url path: `localhost:8081/v1/data/{ip-or-domain}`
### Testing
 - Using Go:
   - Terminal command: `./testing/test.sh`