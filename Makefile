
BINARY_NAME=pokedexcli
BUILD_DIR = ./.out

# Local machine development
build:
	@go build -o ${BUILD_DIR}/${BINARY_NAME}


run: build
	@./${BUILD_DIR}/${BINARY_NAME}

test:
	@go test -v ./...

clean:
	rm -rf $(BUILD_DIR)