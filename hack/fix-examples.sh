#!/bin/bash
set -e

echo "Fixing generated examples..."

# Fix mtu from string to number
find examples-generated -name "*.yaml" -type f -exec sed -i '' 's/mtu: "9000"/mtu: 9000/g' {} \;

# Fix autoCreateIpv4Subnet from string to boolean
find examples-generated -name "*.yaml" -type f -exec sed -i '' 's/autoCreateIpv4Subnet: "false"/autoCreateIpv4Subnet: false/g' {} \;
find examples-generated -name "*.yaml" -type f -exec sed -i '' 's/autoCreateIpv4Subnet: "true"/autoCreateIpv4Subnet: true/g' {} \;

# Fix example-id annotations that are missing the cloudscale prefix
find examples-generated -name "*.yaml" -type f -exec sed -i '' 's|meta.upbound.io/example-id: /v1alpha1/|meta.upbound.io/example-id: cloudscale/v1alpha1/|g' {} \;

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
find examples-generated -name "*.yaml" -type f -exec sed -i '' '/^    count: /d' {} \;

# Fix Terraform variable interpolation in name fields
find examples-generated -name "*.yaml" -type f -exec sed -i '' 's/\${count\.index}/01/g' {} \;

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

