#!/bin/bash
# run as `hack/ext-vendor.sh`

# Define the path to the dependencies.yml file
DEPENDENCIES_FILE="hack/dependencies.yml"
EXTERNAL_DEPENDENCIES_VENDORED_SOURCE="ext-vendor"

# Loop through each entry in the YAML file
yq e '.[]' "$DEPENDENCIES_FILE" -o=json | jq -c '.' | while read -r line; do
  # Extract fields from the JSON
  NAME=$(echo $line | jq -r '.name')
  VERSION=$(echo $line | jq -r '.version')
  REPO=$(echo $line | jq -r '.repo')

  # Construct the directory and download URL
  DIRECTORY="$EXTERNAL_DEPENDENCIES_VENDORED_SOURCE/github.com/$REPO"
  DOWNLOAD_URL="https://github.com/$REPO/archive/refs/tags/$VERSION.tar.gz"

  # Create the directory if it doesn't exist
  mkdir -p "$DIRECTORY"

  # Change to the directory and download the tarball
  pushd "$DIRECTORY" > /dev/null
  echo "Downloading and extracting $NAME from $DOWNLOAD_URL..."
  wget -qO- "$DOWNLOAD_URL" | tar -xz --strip-components=1
  popd > /dev/null
done

echo "All dependencies have been processed."

