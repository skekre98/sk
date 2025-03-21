#!/bin/bash

set -e

echo "Building sk binary..."

# Initialize go module if it doesn't exist
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init sk
    go mod tidy
fi

# Build the binary
echo "Compiling sk..."
go build -o sk

# Check if destination directory exists and is writable
if [ ! -w "/usr/local/bin" ]; then
    echo "Error: /usr/local/bin is not writable. Using sudo to copy the binary."
    sudo cp sk /usr/local/bin/
else
    echo "Copying sk to /usr/local/bin..."
    mv sk /usr/local/bin/
fi

echo "Installation complete! You can now use 'sk' command globally." 