#!/bin/bash
#
#
# Set the directory where you want to search for .env files
search_directory="."

# Find all files with .env extension in the specified directory and its subdirectories
env_files=$(find "$search_directory" -type f -name "*.env")

# Loop through each .env file and run the sops command
for file_name in $env_files; do
    echo "Encrypting $file_name..."
    sops -e -i "$file_name"
done

echo "Encryption completed for all .env files."
