build:
	@go build -o bin/chatterEase

run:build
	@./bin/chatterEase

start:
	@cd webapp && npm start