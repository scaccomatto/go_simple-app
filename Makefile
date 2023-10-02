mocks:
	GO111MODULE=on mockgen -source=./internal/database/database.go -destination=./internal/test/mocks/database/database_mock.go
 
file_tree:
	tree -I vendor

tests:
	go test -v  ./...