# Middle Go & React/TS Fullstack Roadmap

> **Цель:** Выйти на уровень крепкого **Middle Backend / Go-разработчика** (с сильным Fullstack-бэкграундом на React + TypeScript), способного самостоятельно проектировать микросервисы, работать с брокерами сообщений, распределенными блокировками и высокими нагрузками.

---

## 🧭 Обзор архитектурного стека Middle Developer:
- **Язык & Рантайм:** Go (Golang), модель памяти, Escape Analysis, планировщик (GMP), сборщик мусора (GC).
- **Сетевое взаимодействие:** REST (Chi/Gin), **gRPC + Protocol Buffers**, WebSockets (двунаправленная связь).
- **Хранилища данных:** PostgreSQL (ACID, индексы B-Tree/GIN, блокировки, транзакции, профилирование), Redis (кэширование, Pub/Sub, Distributed Locks, Sorted Sets).
- **Асинхронные очереди и брокеры:** **Apache Kafka** или **RabbitMQ** (паттерн Transactional Outbox, консьюмер-группы, идемпотентность).
- **Наблюдаемость (Observability):** Prometheus, Grafana, OpenTelemetry/Jaeger, структурированное логирование (`log/slog`).
- **Тестирование:** Unit (табличные тесты, моки), интеграционные тесты с **Testcontainers-go**, нагрузочное тестирование через **k6**.
- **Frontend:** React 19, TypeScript (строгая типизация), TanStack Query, Tailwind CSS, Canvas API / Pixi.js (для игрового проекта).

---

## Этап 1: Внутренности Go и продвинутая конкурентность (2–3 недели)

Middle-разработчик обязан понимать, как Go работает на уровне железа и операционной системы.

### 1.1. Модель памяти и компилятор
- Стек против кучи: **Escape Analysis** (почему переменная утекает в кучу: `go build -gcflags="-m"`).
- Устройство структур в памяти: выравнивание полей (memory alignment, padding), экономия RAM.
- Устройство интерфейсов: `iface` (с методами) и `eface` (пустой интерфейс), цена динамической диспетчеризации.
- Сборщик мусора (GC): трехцветная маркировка (Tricolor Mark and Sweep), STW (Stop-the-World), переменная `GOGC` и `GOMEMLIMIT`.
- Планировщик Go: модель **GMP** (Goroutine, Machine/OS Thread, Processor), Work Stealing, вытесняющая многозадачность (preemption).

### 1.2. Продвинутые паттерны конкурентности
- Пакет `sync`: `sync.Pool` (переиспользование объектов и снижение нагрузки на GC), `sync.Once`, `sync.Cond`, `atomic` (атомарные операции).
- Паттерны: Worker Pool с динамическим масштабированием, Fan-In / Fan-Out, Pipeline, Rate Limiter (Token Bucket).
- Корректная обработка завершения: Graceful Shutdown с передачей сигналов ОС (`os.Interrupt`, `syscall.SIGTERM`) и `context.WithTimeout`.

---

## Этап 2: Базы данных на уровне Middle (2 недели)

### 2.1. PostgreSQL под нагрузкой
- **Блокировки и конкуренция:**
  - Пессимистическая блокировка: `SELECT ... FOR UPDATE` и `SELECT ... FOR SHARE`.
  - Предотвращение состояния гонки (Race Condition) и взаимных блокировок (**Deadlocks**), порядок захвата ресурсов.
  - Оптимистическая блокировка: версионирование строк (`version INT`).
- **Уровни изоляции транзакций:** Read Committed, Repeatable Read, Serializable. Аномалии: Dirty Read, Non-Repeatable Read, Phantom Read, Serialization Failure.
- **Глубокий анализ производительности:** `EXPLAIN (ANALYZE, BUFFERS)`. Seq Scan, Index Scan, Bitmap Index Scan, CTE Scan.
- **Пул соединений:** Конфигурация `pgxpool.Pool` (MaxConns, MinConns, MaxConnLifetime, MaxConnIdleTime).

### 2.2. Продвинутый Redis
- Структуры данных: Strings, Hashes, Lists, Sets, **Sorted Sets (ZSET)** — алгоритм SkipList под капотом (идеально для игровых лидербордов).
- Паттерны кэширования: Cache-Aside, Write-Through. Проблемы Cache Stampede и Cache Penetration.
- Распределенные блокировки (**Distributed Lock**): концепция `SET key val NX PX` и алгоритм **Redlock**.

---

## Этап 3: Микросервисы, gRPC и брокеры сообщений (3 недели)

### 3.1. gRPC и Protobuf
- Protocol Buffers v3: синтаксис `.proto`, генерация Go-кода через `protoc-gen-go` и `protoc-gen-go-grpc`.
- Типы вызовов: Unary, Server Streaming, Client Streaming, Bidirectional Streaming.
- Перехватчики (Interceptors): логирование, валидация, трассировка, авторизация по токену в метаданных.

### 3.2. Асинхронные брокеры: Kafka / RabbitMQ
- Разница концепций: Smart Broker / Dumb Consumer (RabbitMQ) vs Dumb Broker / Smart Consumer (Kafka).
- **Apache Kafka основы:** Топики, партиции, оффсеты, Consumer Groups, гарантии доставки (At-most-once, At-least-once, Exactly-once semantics).
- Паттерн **Transactional Outbox**: как сохранить данные в Postgres и гарантированно отправить событие в Kafka без распределенных транзакций 2PC.
- **Идемпотентность консьюмеров:** защита от повторной обработки одного и того же сообщения.

---

## Этап 4: Observability, Тестирование и CI/CD (2 недели)

1. **Структурированное логирование:** Стандартный пакет `log/slog` (уровни, форматирование JSON, передача `trace_id`).
2. **Метрики:** Интеграция с **Prometheus** (RED-метод: Rate, Errors, Duration). Метрики задержки (Histogram), счетчики ошибок (Counter), активные соединения (Gauge).
3. **Интеграционные тесты:** Использование библиотеки **Testcontainers-go** для поднятия реальных инстансов PostgreSQL/Redis/Kafka в Docker во время тестов `go test`.
4. **Нагрузочное тестирование:** Написание сценариев на **k6** для проверки пропускной способности (RPS) и времени ответа (p95, p99).
5. **CI/CD:** GitHub Actions (запуск `golangci-lint`, тестов с флагом `-race`, сборка Docker-образов).

---

## Этап 5: Frontend (React 19 + TypeScript) для сложных интерфейсов (2 недели)

- **Строгий TypeScript:** Generic-компоненты, дискриминантные объединения (Discriminated Unions), валидация данных через **Zod**.
- **Стейт-менеджмент:** **TanStack Query** (кэш сервера, оптимистичные обновления) + **Zustand** (клиентское состояние игры/аукциона).
- **Real-Time клиент:** Модульная обертка над WebSocket с автоматическим переподключением, пинг-понгом и очередью неотправленных сообщений.
- **Рендеринг для игры:** Основы работы с **HTML5 Canvas API** или **Pixi.js** (спрайты, отрисовка игрового поля, интерполяция движения).

---

## Этап 6: Боевые пет-проекты (Middle Portfolio)

---

### 🎟️ Проект 1: High-Load Booking & Auction Engine (Fintech / E-Commerce)

**Бизнес-задача:** Сервис мгновенного бронирования билетов и проведения live-аукционов с тысячами одновременных участников.

#### Архитектура и стек:
- **Язык:** Go (Чистая архитектура, DDD-элементы).
- **Базы данных:** PostgreSQL (персистентность) + Redis (кэш, distributed lock).
- **Связь:** REST (HTTP API) + gRPC (внутренние сервисы) + WebSockets (live-ставки).
- **Брокер:** Kafka или RabbitMQ.
- **Фронтенд:** React 19 + TypeScript + Tailwind + TanStack Query.

#### Ключевые фичи:
1. **Защита от овербукинга (Race Condition):**
   - Бронирование места: временный холд на 10 минут с использованием распределенной блокировки в Redis.
   - Финальная оплата: строгая транзакция в PostgreSQL с `SELECT ... FOR UPDATE`.
2. **Live-аукцион через WebSockets:**
   - Канал ставок с мультиплексированием через горутины.
   - Мгновенный broadcast новой максимальной цены всем подключенным зрителям.
3. **Паттерн Transactional Outbox + Worker:**
   - После успешной ставки/покупки запись сохраняется в таблицу `outbox_events` в рамках той же транзакции.
   - Фоновый воркер вычитывает события и пушит их в Kafka для сервиса уведомлений и аналитики.
4. **Инфраструктура и мониторинг:**
   - `docker-compose.yml` (Go App, Postgres, Redis, Kafka, Prometheus, Grafana).
   - Дашборд в Grafana с графиками RPS, длительности транзакций и активных WS-сессий.

---

### 🎮 Проект 2: Real-Time Multiplayer Game Server (GameDev / Low-Latency)

**Бизнес-задача:** Многопользовательская 2D-арена в реальном времени (например, Battle Arena / Space Shooter / Cyber Tactics) с сессионным матчмейкингом и авторитетным сервером.

#### Архитектура и стек:
- **Язык:** Go (Авторитетный сервер — вся физика и правила рассчитываются на бэкенде).
- **Сетевой протокол:** WebSockets с сериализацией данных через **Protocol Buffers** (минимальный вес сетевого пакета).
- **Хранилище:** Redis (Sorted Sets для очередей матчмейкинга и глобального ELO-рейтинга) + PostgreSQL (профили, история матчей).
- **Фронтенд:** React 19 + TypeScript + **HTML5 Canvas / Pixi.js**.

#### Ключевые фичи:
1. **Авторитетный игровой цикл (Authoritative Game Loop):**
   - Каждая комната (Room) — отдельная горутина с тикрейтом 30 или 60 FPS (тик каждые 16.6 или 33.3 мс).
   - Клиенты шлют только команды ввода (направление движения, выстрел).
   - Сервер валидирует действия, обновляет координаты мира и рассылает дельту состояния всем игрокам.
2. **Матчмейкинг на Redis Sorted Sets:**
   - Игроки встают в очередь с текущим MMR/ELO.
   - Воркер ищет соперников с близким рейтингом (`ZRANGEBYSCORE`) и при нахождении создает новую игровую комнату.
3. **Клиентская интерполяция:**
   - Фронтенд на React + Canvas плавно интерполирует координаты между тиками сервера, сглаживая сетевой пинг.
4. **Глобальный лидерборд:**
   - Мгновенный пересчет топа игроков на основе Redis ZSET.
