# Server files

### Use
> create an .env file with the following environment variables

***.env***
```.env
KEY=<apiKey>
HOSTNAME=<host> #you can only see the files from that ip address
```

- EndPoint
    - /api/v1/pfds?key={apiKey} [POST]
    - Response [200 - OK]
        ```json
            {
                "message_type": "message",
                "message": "se guardo correctamente",
                "error": false,
                "data": {
                    "url": "http://192.168.1.6:8080/api/v1/images/TZADhp2ijNp9MjwTXFMx___Captura-de-pantalla-de-2021-08-20-17-48-10.png"
                }
            } 
        ```
    - /api/v1/images?key={apiKey} [POST]
    - Response [200 - OK]
        ```json
            {
                "message_type": "message",
                "message": "se guardo correctamente",
                "error": false,
                "data": {
                    "url": "http://192.168.1.6:8080/api/v1/images/TZADhp2ijNp9MjwTXFMx___Captura-de-pantalla-de-2021-08-20-17-48-10.png"
                }
            } 
        ```
        