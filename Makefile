create_db:
	sudo docker run --name cassandra-db --network some-network -p 9042:9042 -d cassandra:latest
start_db:
	sudo docker start cassandra-db
db_shell:
	sudo docker run -it --rm --network some-network cassandra:latest cqlsh cassandra-db
check_install:
	which swagger || GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_install
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models
test:
	GO111MODULE=on go test ./... -v
build:
	GO111MODULE=on go mod download
run:
	GO111MODULE=on go run main.go
