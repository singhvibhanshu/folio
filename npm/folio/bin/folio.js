#!/usr/bin/env node
'use strict';

// Launcher for the folio CLI.
//
// folio is a Go program shipped as prebuilt, per-platform binaries. The matching
// binary lives in an optionalDependency that npm installs only for the user's
// OS/CPU. This script resolves that binary and execs it, forwarding all args.
// Nothing here touches the network — installation is the only online step.

const { spawnSync } = require('child_process');

const PLATFORMS = {
  'darwin-arm64': 'darwin-arm64',
  'darwin-x64': 'darwin-x64',
  'linux-x64': 'linux-x64',
  'linux-arm64': 'linux-arm64',
  'win32-x64': 'win32-x64',
};

const key = `${process.platform}-${process.arch}`;
const suffix = PLATFORMS[key];

if (!suffix) {
  console.error(`folio: unsupported platform "${key}".`);
  console.error('Supported: macOS (arm64/x64), Linux (x64/arm64), Windows (x64).');
  process.exit(1);
}

const pkg = `@singhvibhanshu/folio-${suffix}`;
const binName = process.platform === 'win32' ? 'folio.exe' : 'folio';

let binPath;
try {
  binPath = require.resolve(`${pkg}/bin/${binName}`);
} catch (_) {
  console.error(`folio: prebuilt binary for "${key}" was not found.`);
  console.error(`Expected the optional dependency "${pkg}" to be installed.`);
  console.error('Reinstall with: npm install @singhvibhanshu/folio');
  process.exit(1);
}

const result = spawnSync(binPath, process.argv.slice(2), { stdio: 'inherit' });

if (result.error) {
  console.error('folio: failed to start binary:', result.error.message);
  process.exit(1);
}
process.exit(result.status === null ? 1 : result.status);
