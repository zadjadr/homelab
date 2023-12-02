#!/bin/bash
#
#
# Set the directory where you want to search for .env files
search_directory="."

# Find all files with .env extension in the specified directory and its subdirectories
env_files=$(find "$search_directory" -type f -name "*.env" -not -path "$search_directory/tools/wireguard/config/*" 2>/dev/null)

# Loop through each .env file and run the sops command
for file_name in $env_files; do
    if sops -d "$file_name" > /dev/null 2>&1; then
        echo "✅ $file_name is already encrypted."
    else
        echo "🔐 Encrypting $file_name..."
        sops -e -i "$file_name"
    fi
done

echo "Encryption completed for all .env files."
