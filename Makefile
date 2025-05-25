CONFIG ?= ./scripts_config/bluetooth_ctl.yaml

.PHONY: all build run clean test

# Build the binary
build:
	@go build -o xofikit ./cmd/xofikit/main.go

# Run with provided config or default
run:
	./xofikit run $(CONFIG)

# Clean up binary
clean:
	rm -f xofikit

install:
	@echo "Downloading dependencies"
	@go mod tidy
	@echo "[INFO] Building xofikit..."
	@make build
	@echo "[INFO] Setting executable permissions..."
	@chmod +x ./xofikit
	@echo "[INFO] Moving binary to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	@cp ./xofikit ~/.local/bin/
	@sudo mv ./xofikit /usr/bin/

	@echo "[INFO] Copying scripts_config to ~/.config/xofikit..."
	@mkdir -p ~/.config/xofikit
	@cp -r ./scripts_config ~/.config/xofikit/

	@echo "[INFO] Installation completed successfully."

uninstall:
	@rm -f ~/.local/bin/xofikit
	@rm -rf ~/.config/xofikit
	@echo "xofikit uninstalled."

# Build, run, and clean in one go
test: build run clean
