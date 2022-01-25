ITS_OVER_NINE_THOUSEEEEEND = 9001
ENVS = export $(cat .env | xargs)
GOCMD = go
GOTEST = $(GOCMD) test
GO_PRJ_MAIN = cmd/badge/main.go
BIN_NAME = badge-svc

all: run

build:
	go build -o=$(BIN_NAME) $(GO_PRJ_MAIN)

run:
	PORT=$(ITS_OVER_NINE_THOUSEEEEEND) DATABASE_URL=postgres://cuely:cuely@localhost:5432/cuely?sslmode=disable JWT_SECRET=SUPERSECRET $(GOCMD) run $(GO_PRJ_MAIN)

test:
	$(GOTEST) ./...
