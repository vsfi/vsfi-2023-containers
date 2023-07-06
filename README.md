# vsfi-2023-containers
## Architecure
Services
* svc - service for register tickets
* admin - service for submit tickets
* db - database for store tickets

## Run 
Run with docker compose
```
docker compose up -d 

[+] Running 4/4
 ⠿ Network vsfi-2023-containers_default  Created        0.1s
 ⠿ Container db                          Healthy        0.7s
 ⠿ Container svc                         Started        0.6s
 ⠿ Container admin                       Started        0.6s
```

## Check
Check for services is works
```
NAME                COMMAND                  SERVICE  STATUS              PORTS
admin               "python3 app.py"         admin    running             0.0.0.0:5000->5000/tcp, :::5000->5000/tcp
db                  "/cockroach/cockroac…"   db       running (healthy)   8080/tcp, 26257/tcp
svc                 "/app/beer"              svc      running             0.0.0.0:8080->8080/tcp, :::8080->8080/tcp
```
