#!/bin/bash
set -e

echo "Fixing generated examples..."

# Detect OS and set sed in-place flag accordingly
# macOS/BSD sed requires -i '' while GNU/Linux sed requires -i
SED_INPLACE="-i"
if [[ "$OSTYPE" == "darwin"* ]]; then
    SED_INPLACE="-i ''"
fi

# Function to run sed in-place compatible with both Linux and macOS
sed_inplace() {
    local pattern="$1"
    local file="$2"
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' "$pattern" "$file"
    else
        sed -i "$pattern" "$file"
    fi
}

# Fix mtu from string to number
for file in $(find examples-generated -name "*.yaml" -type f); do
    sed_inplace 's/mtu: "9000"/mtu: 9000/g' "$file"
done

# Fix autoCreateIpv4Subnet from string to boolean
for file in $(find examples-generated -name "*.yaml" -type f); do
    sed_inplace 's/autoCreateIpv4Subnet: "false"/autoCreateIpv4Subnet: false/g' "$file"
    sed_inplace 's/autoCreateIpv4Subnet: "true"/autoCreateIpv4Subnet: true/g' "$file"
done

# Fix example-id annotations that are missing the cloudscale prefix
for file in $(find examples-generated -name "*.yaml" -type f); do
    sed_inplace 's|meta.upbound.io/example-id: /v1alpha1/|meta.upbound.io/example-id: cloudscale/v1alpha1/|g' "$file"
done

# Remove timeouts blocks (these are Terraform-specific)
# This is a bit tricky with sed, so we'll use a more robust approach
for file in examples-generated/cluster/cloudscale/v1alpha1/*.yaml examples-generated/namespaced/cloudscale/v1alpha1/*.yaml; do
    if [ -f "$file" ]; then
        # Remove timeouts blocks using awk
        awk '
            /^    timeouts:/ { skip=1; next }
            skip && /^    [a-z]/ { skip=0 }
            skip && /^---/ { skip=0 }
            !skip { print }
        ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
    fi
done

# Remove count fields (Terraform meta-argument not applicable in Crossplane)
for file in $(find examples-generated -name "*.yaml" -type f); do
    sed_inplace '/^    count: /d' "$file"
done

# Fix Terraform variable interpolation in name fields
for file in $(find examples-generated -name "*.yaml" -type f); do
    sed_inplace 's/\${count\.index}/01/g' "$file"
done

# Remove duplicate SSH keys using awk to track seen keys per file
for file in examples-generated/cluster/cloudscale/v1alpha1/*.yaml examples-generated/namespaced/cloudscale/v1alpha1/*.yaml; do
    if [ -f "$file" ]; then
        awk '
            /^    - (ssh-|ecdsa-)/ {
                if (!seen[$0]++) {
                    print
                }
                next
            }
            /^[^ ]/ || /^  [^ ]/ || /^    [^-]/ {
                delete seen
            }
            { print }
        ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
    fi
done

echo "Examples fixed!"

