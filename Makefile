sqlup:
	cd ./sso
	migrate -path sso/internal/migrations -database "postgresql://root:12345@localhost:5432/root?sslmode=disable" -verbose up
sqldown:
	cd ./sso
	migrate -path sso/internal/migrations -database "postgresql://root:12345@localhost:5432/root?sslmode=disable" -verbose down
grpc:
	cd ./protos
	protoc -I . ./sso.proto --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative
.PHONY: sqlup sqldown grpc