# folio

[![npm version](https://img.shields.io/npm/v/@singhvibhanshu/folio.svg)](https://www.npmjs.com/package/@singhvibhanshu/folio)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![platforms](https://img.shields.io/badge/platforms-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg)](#install)

**A privacy-first, fully offline PDF toolkit.** Merge, split, compress, protect,
watermark, reorder pages and more, all on your own machine. **Nothing is ever
uploaded.** No servers, no accounts, no internet. Your documents never leave
your computer.

> Like the online "I♥PDF"-style tools, but your files stay private. folio is a
> single self-contained binary (no runtime, no system libraries), the offline
> sibling of [imago](https://github.com/singhvibhanshu/imago), which does the
> same for images.

## Install

```bash
# npm (global)
npm install -g @singhvibhanshu/folio
folio --help

# or run without installing
npx @singhvibhanshu/folio --help

# Bun works too (same npm registry)
bunx @singhvibhanshu/folio --help

# Go developers
go install github.com/singhvibhanshu/folio@latest
```

Installation downloads only the prebuilt binary matching your OS/CPU. After
that, folio runs **100% offline**.

## Commands

| Command | What it does | Example |
|---|---|---|
| `merge` | Combine PDFs into one | `folio merge a.pdf b.pdf -o out.pdf` |
| `split` | Split into per-page (or N-page) files | `folio split big.pdf --span 1` |
| `compress` | Shrink file size, losslessly | `folio compress report.pdf` |
| `rotate` | Rotate pages by 90/180/270° | `folio rotate scan.pdf --angle 90` |
| `protect` | Password-protect (AES-256) | `folio protect taxes.pdf -p secret` |
| `unlock` | Remove a known password | `folio unlock taxes.pdf -p secret` |
| `watermark` | Stamp text across pages | `folio watermark doc.pdf -t DRAFT` |
| `pagenum` | Add "n of N" page numbers | `folio pagenum thesis.pdf` |
| `extract` | Keep / reorder selected pages | `folio extract in.pdf --pages 3,1,2` |
| `remove` | Delete selected pages | `folio remove in.pdf --pages 2,4` |
| `frompics` | Build a PDF from images | `folio frompics *.jpg -o album.pdf` |
| `images` | Extract embedded images | `folio images scan.pdf` |
| `nup` | N pages per sheet (handouts) | `folio nup slides.pdf -n 4` |
| `crop` | Trim page margins | `folio crop in.pdf --box "10% 10% 80% 80%"` |
| `info` | Show pages, size, metadata | `folio info report.pdf` |

Run `folio <command> --help` for every flag.

### Handy details

- **Batch a folder.** Most single-file commands accept a directory and process
  every `*.pdf` inside it: `folio compress ./invoices -o ./out`.
- **Page selection.** `--pages` takes ranges and lists like `1-3,5,8-` and, for
  `extract`, preserves the order you give (so it also reorders).
- **Output naming.** With no `--out`, folio writes next to the input with a
  descriptive suffix (`report.min.pdf`, `scan.rotated.pdf`, …). Point `--out` at
  a file to name it, or at a directory to drop outputs there.

## Why offline matters

Online PDF tools require you to upload your documents (contracts, IDs, tax
forms, medical records) to someone else's server. folio does the same
operations locally, so sensitive files never leave your machine. As a bonus,
re-saving a PDF (compress, etc.) discards hidden metadata along the way.

## Scope & honest limitations

folio covers the operations that can be done **purely offline in a single Go
binary** via [pdfcpu](https://github.com/pdfcpu/pdfcpu). A few tools the online
services offer are deliberately **not** included, because they require
heavyweight engines (LibreOffice, Tesseract, a PDF rasterizer) that would break
the no-internet, single-binary promise:

- PDF ⇄ Word / Excel / PowerPoint conversion
- OCR (making scanned PDFs searchable)
- Rasterizing whole pages to JPG/PNG (folio's `images` extracts *embedded*
  images, it does not screenshot pages)

## Build from source

```bash
git clone https://github.com/singhvibhanshu/folio.git
cd folio
go build -o folio .
./folio --help
```

## Maintainer: releasing

Binaries are cross-compiled and published from one place.

```bash
# 1. bump the version in npm/folio/package.json
# 2. tag and push: GitHub Actions builds all platforms and publishes to npm
git tag v0.1.1 && git push --tags
```

CI needs an `NPM_TOKEN` repository secret (a granular npm token with read/write
to the `@singhvibhanshu` scope). To publish manually instead:

```bash
bash scripts/build-npm.sh      # cross-compile all 5 platforms
bash scripts/publish-npm.sh    # publish platform packages, then the main one
```

## License

MIT © 2026 Vibhanshu Singh
