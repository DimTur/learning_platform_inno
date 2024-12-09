**Данный репозиторий собрает воедино все сервисы в рамках выпускной работы**

**Для того, чтобы развернуть локально необходимо сделать следующие шаги:**

1. Собрать докер файлы из каждого сервиса. Команды предоставлены в каждом.
2. Перейти в папку 'k8s'
3. Создать kind кластер:

        kind create cluster --config=kind.yaml

4. Добавить образы в kind:

        kind load docker-image job-lp-migrator:1.0.0
        kind load docker-image app-lp:1.0.0
        kind load docker-image app-sso:1.0.0
        kind load docker-image job-queue-manager:1.0.0
        kind load docker-image app-notification:1.0.0
        kind load docker-image app-api-gateway:1.0.0

5. Генерируем и добавляем ключи для генерации токенов:
        
        openssl genpkey -algorithm Ed25519 -out private_key.pem
        openssl pkey -in private_key.pem -pubout -out public_key.pem

        kubectl create secret generic sso-keys \
        --from-file=private_key.pem=./private_key.pem \
        --from-file=public_key.pem=./public_key.pem

6. Последовательность применения манифестов:

            kubectl apply -f ./configmaps
            kubectl apply -f ./pvs
            kubectl apply -f ./services
            kubectl apply -f ./jobs
            kubectl apply -f ./deployments/dependencies

    - Для сервиса notification необходимо в конфиг мапах вставить актуальный токен вашего бота.

    - Необходимо создать пользователя для монги:
            
            kubectl exec deployment/mongo-deployment -it -- /bin/bash
            mongosh
            use admin
            db.createUser({user: "test", pwd: "test", roles: [{role: "root", db: "admin"}]})

    - Применяем оставшиеся манифесты:

            kubectl apply -f ./deployments

7. Приложение доступно на порту 30000. Свагер по адресу: http://localhost:30000/swagger/index.html