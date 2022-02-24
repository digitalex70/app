BINARY_NAME=app

build:
	@echo "Building Celeritas..."
	@go build -o ./build/bin/${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	@./build/bin/${BINARY_NAME} &
	@echo "Celeritas started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ./build/bin/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./build/bin/$${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start