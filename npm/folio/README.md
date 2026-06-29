# @singhvibhanshu/folio

**A privacy-first, fully offline PDF toolkit.** Merge, split, compress, protect,
watermark, reorder pages and more — entirely on your machine. Nothing is
uploaded; your documents never leave your computer.

```bash
# install globally
npm install -g @singhvibhanshu/folio
folio --help

# or run without installing
npx @singhvibhanshu/folio --help
bunx @singhvibhanshu/folio --help
```

Only installation uses the network (to fetch the prebuilt binary for your
platform). Every operation afterwards runs 100% offline.

## Commands

| Command | What it does |
|---|---|
| `merge` | Combine PDFs into one |
| `split` | Split into per-page (or N-page) files |
| `compress` | Shrink file size, losslessly |
| `rotate` | Rotate pages by 90/180/270° |
| `protect` / `unlock` | Add / remove an AES-256 password |
| `watermark` | Stamp text across pages |
| `pagenum` | Add "n of N" page numbers |
| `extract` / `remove` | Keep+reorder / delete selected pages |
| `frompics` | Build a PDF from images |
| `images` | Extract embedded images |
| `nup` | N pages per sheet |
| `crop` | Trim page margins |
| `info` | Show pages, size, metadata |

```bash
folio merge a.pdf b.pdf -o combined.pdf
folio compress report.pdf
folio protect taxes.pdf --password secret
folio extract slides.pdf --pages 3,1,2
```

Run `folio <command> --help` for all flags.

## Notes

folio is a single self-contained binary (Go) and stays fully offline. It does
**not** include Office-format conversion, OCR, or whole-page rasterization,
since those need heavyweight external engines that would break the
single-binary, no-internet design.

Source & issues: https://github.com/singhvibhanshu/folio

MIT © 2026 Vibhanshu Singh
