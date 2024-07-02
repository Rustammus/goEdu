## Любые замечания и советы приветствуются
Это мой пет проект. Разрабатываю с целью получить опыт разработки.

На данный момент проект очень сырой, фунционала нет, пока только базоый CRUD, меня интересует оценка архитектуры, на сколько код читаем, и прочие замечания.
Пока реализованны: аутентификация (refresh token не реализован), базовый CRUD над моделью User.

Коротко о сервисе, который я хочу разработать. План так же сырой. \
Это онлайн сервис для обработки медиа (смена разрешения, конвертация в другие форматы и многое другое, что предоствляет ffmpeg)
Сервис будет состоять из частей:
 - Бекенд сервер, взаимодействущий с пользователем, отправляет задачи обработки видео на брокер сообщений (RabbitMQ) и котролирует выполнение
 - Кластер k8s, обрабатывющий задачи. 1 мастер нода и 1-2 воркер ноды. На воркер ноде множество контейнеров с ffmpeg, читают сообщения с задачами
 - S3 хранилище для хранения медиа, пользователь будет загружать медиа сразу в хранилище, по сгенерированной временной ссылке. ffmpeg контейнеры так же будут взаимодествовать с хранилищем
 - Возможно frontend, меня не очень интересует, но взаимодествовать через сваггер не очень.

Запуск: \
БД - Postgresql \
Миграции выполнять с помощью инструмента goose \
$ go install github.com/pressly/goose/v3/cmd/goose@latest \
$ goose postgres "user=postgres dbname=postgres sslmode=disable" reset \
Либо просто выполнить sql команды в папке migration \

Swagger доступен по localhost:8082/swagger/index.html

