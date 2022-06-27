# ip-data
## Building & Running:
### AWS Lambda
 - Configuration
   - Add an environment variable with key=DOMAIN_API_KEY and value=your-api-key-to-api.ip2whois.com [Docs](https://www.ip2whois.com/developers-api)
 - Build
   - Run `./buildLambda.sh` and upload to aws lambda function

### Local
 - Configuration
   - `$export DOMAIN_API_KEY="your-api-key-to-api.ip2whois.com"` This is used to query data for domain names. [Docs](https://www.ip2whois.com/developers-api)
   - `$export IP_DATA_SERVER_PORT=":80"` Default server port is 8081
 - Using Go: 
   - [Install Go](https://golang.org/doc/install)
   - Terminal command: ``go run main.go``
   - Url path: `localhost:8081/v1/data/{ip-or-domain}`
### Testing
 - Using Go:
   - Terminal command: `./testing/test.sh`

 - Using curl:
   - IP address Example:
      ```
      curl --location --request GET 'https://idzvlkcor8.execute-api.us-west-2.amazonaws.com/default/GetIpData/2601:680:8300:933:dd7:1e1a:c2f5:203d' \
      --header 'Accept: application/json'
      ```
   - Domain Example:
      ```
      curl --location --request GET 'https://idzvlkcor8.execute-api.us-west-2.amazonaws.com/default/GetIpData/google.com' \
      --header 'Accept: application/json'
      ```