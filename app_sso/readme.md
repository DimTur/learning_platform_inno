**Before start you need to apply migrations and generate keys**

    go run ./cmd/migrator apply --config=./config/config.yaml --migrationsPath=./migrations

    openssl genpkey -algorithm Ed25519 -out private_key.pem
    openssl pkey -in private_key.pem -pubout -out public_key.pem

**Start service**

    go run cmd/main.go serve --config=./config/config.yaml


docker compose up -d
docker compose down -v