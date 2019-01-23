$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o ./bin/alertmangetokafka ./main.go