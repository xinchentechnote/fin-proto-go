
# 运行测试
test:
	gotest -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html