#!/usr/bin/env bash
#
# Publish folio to npm. The 5 platform packages must go up BEFORE the main
# package (which lists them as optionalDependencies), so the order matters.
#
#   bash scripts/publish-npm.sh
#
# Interactively (a real terminal): you'll be asked for a one-time 2FA code
# before each package. Leave it blank if your account/token doesn't need one.
# In CI (no terminal): it publishes non-interactively using NODE_AUTH_TOKEN.
#
set -euo pipefail
cd "$(dirname "$0")/.."

if [ ! -d npm/platforms ]; then
  echo "No npm/platforms found. Run: bash scripts/build-npm.sh first." >&2
  exit 1
fi

publish_dir() {
  local dir="$1"
  local name version
  name=$(node -e "console.log(require('./$dir/package.json').name)")
  version=$(node -e "console.log(require('./$dir/package.json').version)")
  # Skip if this exact name@version is already on npm, so re-runs (e.g. a
  # re-pushed tag) succeed instead of failing with "cannot publish over
  # previously published version".
  if npm view "$name@$version" version >/dev/null 2>&1; then
    echo ">> already published, skipping $name@$version"
    return 0
  fi
  echo ">> publishing $name"
  if [ -t 0 ]; then
    read -r -p "   2FA OTP (blank if none): " otp
    if [ -n "$otp" ]; then
      ( cd "$dir" && npm publish --access public --otp="$otp" )
    else
      ( cd "$dir" && npm publish --access public )
    fi
  else
    ( cd "$dir" && npm publish --access public )
  fi
}

# Platform packages first...
for d in npm/platforms/*/; do
  publish_dir "${d%/}"
done

# ...then the main package.
publish_dir "npm/folio"

echo "All packages published."
