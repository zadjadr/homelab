#!/bin/bash
#
#
# Set the directory where you want to search for .env files
search_directory="."

# Find all files with .env extension in the specified directory and its subdirectories
env_files=$(find "$search_directory" -type f -name "*.env" -not -path "$search_directory/tools/wireguard/config/*" 2>/dev/null)

# Loop through each .env file and run the sops command
for file_name in $env_files; do
    decrypted_file="$file_name.dec"

    rm -f $decrypted_file
    echo "Decrypting $file_name into $decrypted_file..."
    # Run the sops command and save the decrypted content to the new file
    sops -d "$file_name" > "$decrypted_file"
    chmod 440 "$decrypted_file"
done

echo "Decryption completed for all .env files."
