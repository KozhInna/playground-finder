# Playground Finder: plan

Stages in the order to be build.

| # | Stage | Why here |
|---|---|---|
| 1 | **Skeleton**: module, folders, config, DB connection, `/health` | Everything else plugs into it. |
| 2 | **HRI client and sync**: fetch, filter 475/389, upsert, mark removed | Without data, nothing else can be tested. |
| 3 | **Read API**: `GET /playgrounds` with the `has_building` filter and distance sort | The core of the product, and pure Go logic that is easy to unit test. |
| 4 | **Minimal map frontend**: one Leaflet and OpenStreetMap page served by Gin | Gives a visible result early, and forces the frontend decision. |
| 5 | **Auth**: Google sign-in, users table, auth middleware | Photos and reviews depend on it. |
| 6 | **Photos and reviews**: CRUD with ownership checks | Needs auth first, and the file-storage decision. |
| 7 | **Polish**: central error handling, Swagger, README with Limitations, Dockerfile, Makefile | |
| 8 | **Later stages**: scheduled sync, Compose and GitHub Action, integration tests, PostGIS | |


## Task breakdown

Stages 2–8 are not split yet. We split one stage at a time, just before
starting it.

## Stage 1: Skeleton (3–4 sessions)

**Task 1: Module, folders and config.**
`go mod init`, the `cmd/server` and `internal/` folders, `.gitignore`,
`.env.sample`, and the `internal/config` package that loads the DB URL and
port from the environment.
*Done when:* `go run ./cmd/server` prints the port and DB URL read from
`.env`; removing a variable gives a clear error message; `.env` is ignored by
git and `.env.sample` is committed.

**Task 2: DB connection.**
The `internal/db` package opens a pgx connection pool and pings it. The owner
types the first query: `SELECT count(*) FROM playgrounds`.
*Done when:* the program prints the count (0 is fine, as the table stays
empty until the stage 2 sync), and a wrong password in `.env` gives a
readable error instead of a crash.

**Task 3: Gin server and /health.**
`main.go` wires config, DB and router together. `/health` pings the DB.
*Done when:* `curl localhost:8080/health` returns 200 with the DB status OK,
and after `docker stop playground-db` it returns 503.

**Task 4: Makefile and first test.**
`make dev` and `make test`, plus `config_test.go` next to `config.go`.
*Done when:* `make dev` starts the server, `make test` passes, and the test
covers the "missing variable" case.
