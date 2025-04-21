#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

echo "Starting supervisor_start.sh script..."

# --- Wait for Ollama Server ---
echo "Waiting for Ollama server to be ready at http://localhost:11434..."
attempt=0
max_attempts=60 # Wait up to 5 minutes (60 * 5 seconds)

# Use 'curl --silent --fail' which returns non-zero on failure (like connection refused or 404)
# The '>' redirects stdout (the actual page content) to /dev/null
while ! curl --silent --fail http://localhost:11434 > /dev/null; do
  attempt=$((attempt + 1))
  if [ "$attempt" -ge "$max_attempts" ]; then
    echo "Error: Ollama server did not become ready after $max_attempts attempts."
    exit 1
  fi
  echo "Ollama not ready, waiting 5 seconds... (Attempt $attempt/$max_attempts)"
  sleep 5
done
echo "Ollama server is ready."

# --- Pull the required model ---
# Check if the model already exists locally before pulling
# Note: 'ollama list' output format might change, this is a basic check
MODEL_NAME="granite-embedding"
echo "Checking if model '$MODEL_NAME' exists..."
if ! ollama list | grep -q "^$MODEL_NAME"; then
  echo "Model '$MODEL_NAME' not found locally. Pulling..."
  ollama pull "$MODEL_NAME"
  echo "Model '$MODEL_NAME' pulled successfully."
else
  echo "Model '$MODEL_NAME' already exists locally."
fi

# --- Start the Go Application ---
echo "Starting Go application (/app/server)..."
cd /app && exec /app/server
