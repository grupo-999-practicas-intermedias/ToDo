# prac-intermedias-g999

## Para crear un workspace
```bash
# Correr en la carpeta del proyecto
yarn init 

# crear directorios
mkdir server
mkdir client

```

## Instalar e iniciar Redis en Docker
```bash
# instalar iniciar contenedor
docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
# Visitar puerto 8001
http://localhost:8001/redis-stack/browser
```



## Iniciar backend - golang
```bash
# Correr en la carpeta del proyecto
cd server
go mod init <nombre>
# instalar dependencias
go get github.com/gofiber/fiber/v2
# correr 
go run server.go
```

## Iniciar frontend - sveltekit
```bash
# Correr en la carpeta del proyecto
cd client
# instalar sveltekit
npm create svelte@latest
# instalar dependencias
npm install
# correr
npm run dev -- --open
```

## unir todo en el workspace
```json
  "scripts": {
    "dev:server":" cd ./server && go run server.go"
  }
```
