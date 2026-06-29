#!/usr/bin/env bash
#
# Cross-compile folio for every supported platform and lay out the npm packages.
# Run from anywhere; paths are resolved relative to the repo root.
#
#   bash scripts/build-npm.sh
#
set -euo pipefail
cd "$(dirname "$0")/.."

SCOPE="@singhvibhanshu"
NAME="folio"
VERSION=$(node -e "console.log(require('./npm/folio/package.json').version)")

echo "Building $NAME v$VERSION for all platforms..."

# entry format: GOOS GOARCH pkg-suffix npm-os npm-cpu binary-name
PLATFORMS=(
  "darwin  arm64 darwin-arm64 darwin arm64 folio"
  "darwin  amd64 darwin-x64   darwin x64   folio"
  "linux   amd64 linux-x64    linux  x64   folio"
  "linux   arm64 linux-arm64  linux  arm64 folio"
  "windows amd64 win32-x64    win32  x64   folio.exe"
)

rm -rf npm/platforms

for entry in "${PLATFORMS[@]}"; do
  read -r GOOS GOARCH SUFFIX NPMOS NPMCPU BIN <<< "$entry"
  PKGDIR="npm/platforms/${NAME}-${SUFFIX}"
  mkdir -p "$PKGDIR/bin"

  echo "  - ${GOOS}/${GOARCH} -> ${PKGDIR}/bin/${BIN}"
  CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" \
    go build -trimpath -ldflags "-s -w" -o "${PKGDIR}/bin/${BIN}" .

  cat > "${PKGDIR}/package.json" <<EOF
{
  "name": "${SCOPE}/${NAME}-${SUFFIX}",
  "version": "${VERSION}",
  "description": "folio prebuilt binary for ${SUFFIX}.",
  "os": ["${NPMOS}"],
  "cpu": ["${NPMCPU}"],
  "files": ["bin/"],
  "license": "MIT"
}
EOF
done

# Keep the main package's optionalDependencies pinned to the same version.
node -e '
  const fs = require("fs");
  const p = "./npm/folio/package.json";
  const j = JSON.parse(fs.readFileSync(p, "utf8"));
  for (const k of Object.keys(j.optionalDependencies || {})) j.optionalDependencies[k] = j.version;
  fs.writeFileSync(p, JSON.stringify(j, null, 2) + "\n");
  console.log("Synced optionalDependencies -> " + j.version);
'

echo "Done. Platform packages are in npm/platforms/."
