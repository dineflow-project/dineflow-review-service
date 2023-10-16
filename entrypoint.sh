# wait-for "localhost:8080" -- "$@"
wait-for "${MYSQL_HOST}:${PORT}" -- "$@"
# Watch your .go files and invoke go build if the files changed.
CompileDaemon --build="go build -o main main.go"  --command=./main