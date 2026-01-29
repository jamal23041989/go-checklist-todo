#



## Architecture
```
/my-checklist-project
│
├── /cmd
│   ├── /api-gateway         # Сервис 1: API (HTTP -> Kafka/gRPC)
│   │   ├── main.go          # Точка сборки API: DI, запуск http.Server
│   │   └── Dockerfile
│   ├── /db-core             # Сервис 2: Core (gRPC -> Postgres/Redis)
│   │   ├── main.go          # Точка сборки Core: DI, запуск grpc.Server
│   │   └── Dockerfile
│   └── /worker              # Сервис 3: Logger (Kafka Consumer)
│       ├── main.go          # Точка сборки Worker: запуск Kafka Consumer
│       └── Dockerfile
│
├── /internal
│   ├── /core                # --- Слой Shared Kernel ---
│   │   ├── /config          # Загрузка env (через cleanenv или viper)
│   │   ├── /logger          # Инициализация slog (интерфейс для всего приложения)
│   │   ├── /apperr          # Типы ошибок: sentinel errors (ErrNotFound) и Error Wrapper
│   │   ├── /middleware      # HTTP/gRPC перехватчики (Recovery, Logging, Auth)
│   │   ├── /tools           # Маленькие утилиты (валидация uuid, парсинг дат)
│   │   └── /domains         # Business Entities: Task struct, Event struct (чистый Go)
│   │
│   ├── /infrastructure      # --- Слой реализации адаптеров ---
│   │   ├── /postgres        # pgxpool.New(), методы для транзакций
│   │   ├── /redis           # redis.NewClient(), методы Get/Set (технические)
│   │   └── /kafka           # Настройка Reader/Writer, логика ретраев
│   │
│   ├── /features            # --- Слой фич (Vertical Slices) ---
│   │   ├── /api_tasks       # [Сервис Gateway]
│   │   │   ├── /handler     # http_handler.go (декодинг JSON, вызов сервиса)
│   │   │   └── /service     # task_service.go (вызов gRPC клиента + Kafka producer)
│   │   │
│   │   ├── /db_tasks        # [Сервис DB-Core]
│   │   │   ├── /handler     # grpc_handler.go (маппинг proto -> domain)
│   │   │   ├── /service     # logic.go (бизнес-правила, логика кеширования в Redis)
│   │   │   └── /repository  # pg_repo.go (SQL запросы через pgx)
│   │   │
│   │   └── /worker_tasks    # [Сервис Worker]
│   │       ├── /handler     # kafka_handler.go (слушатель топика, анмаршалинг)
│   │       └── /service     # process_logic.go (запись в файл/лог)
│   │
│   └── /generated           # --- gRPC код из proto (Сгенерированный код) ---
│       └── /checklist_pb    # *.pb.go и *_grpc.pb.go (результат protoc)
│
├── /proto                   # Контракты: checklist.proto
├── /migrations              # SQL миграции (00001_create_tasks_table.sql)
├── /tests                   # /integration (тесты БД), /unit (тесты сервисов)
├── Makefile                 # Команды: proto, migrate, build, run
├── docker-compose.yml       # Описание: postgres, redis, kafka, zookeeper, apps
└── go.mod                   # Один модуль на весь проект
```
