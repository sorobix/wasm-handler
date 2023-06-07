package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func BuildWasmFromMainRS(cargoToml string, mainRSContent string) RedisObject {
	result := RedisObject{
		Success: false,
	}

	tmpDir, err := os.MkdirTemp("", "rust-project")
	if err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to create temporary directory: %v", err))
		return result
	}
	defer os.RemoveAll(tmpDir)

	if err := os.Chdir(tmpDir); err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to change directory: %v", err))
		return result
	}

	projectDir := filepath.Join(tmpDir, "sorobix_temp")

	if err := exec.Command("cargo", "new", "--lib", projectDir).Run(); err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to create rust project: %v", err))
		return result
	}

	if err := os.WriteFile(filepath.Join(projectDir, "Cargo.toml"), []byte(cargoToml), 0644); err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to write cargo.toml file: %v", err))
		return result
	}

	if err := os.WriteFile(filepath.Join(projectDir, "src", "lib.rs"), []byte(mainRSContent), 0644); err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to write lib.rs file: %v", err))
		return result
	}

	if err := os.Chdir(projectDir); err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to change directory: %v", err))
		return result
	}

	cmd := exec.Command("cargo", "build", "--target", "wasm32-unknown-unknown", "--release")
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		errorMessage := strings.TrimSpace(string(stderr))
		result.Message = base64Encoder(fmt.Sprintf("failed to compile rust project: %v", errorMessage))
		return result
	}

	wasmFilePath := filepath.Join(projectDir, "target", "wasm32-unknown-unknown", "release", "sorobix_temp.wasm")
	if _, err := os.Stat(wasmFilePath); os.IsNotExist(err) {
		result.Message = base64Encoder(fmt.Sprintf("wasm file does not exist: %v", err))
		return result
	}

	wasmData, err := os.ReadFile(wasmFilePath)
	if err != nil {
		result.Message = base64Encoder(fmt.Sprintf("failed to read wasm file: %v", err))
		return result
	}

	b64EncodedWasmFile := base64Encoder(string(wasmData))

	result.Success = true
	result.Message = base64Encoder("Compilation successful!")
	result.Wasm = b64EncodedWasmFile

	return result
}
