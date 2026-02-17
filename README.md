```shell
docker compose up -d
k6 run k6/load.js
```

Open http://localhost:3000, use admin:admin to get access.

### Note

If you use mac and docker desktop, make sure "Enable host networking" is on under Settings -> Resources -> Network.