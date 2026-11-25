# Changelog

Semua perubahan penting didokumentasikan di sini. Format mengikuti
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), versi mengikuti
[SemVer](https://semver.org/).

Repo ini adalah **starter aplikasi + tooling** ekosistem donisgo: ia
mengonsumsi `github.com/rolldone/donisgo` sebagai dependency yang di-pin di
`go.mod`, dan tidak dipublikasikan untuk di-`go get` (konsumsi lewat
clone/fork — GitHub *Use this template*).

## [Unreleased]

### Added

- **SPEC-005 fase 2** — terima aset aplikasi dari repo `donisgo`
  (commit sumber `dfdd4f5`, 2026-11-25) sebagai *fresh copy* tanpa git history
  antar-repo:
  - `plugins/pages/` — plugin konten "pages" (SPA serving `/app`, redirect
    `/spa` → `/app`, proxy homepage via `HOMEPAGE_URL`, wildcard `NoRoute`)
    beserta testnya. 349 baris (`plugin.go` 159 + `plugin_test.go` 190).
  - `cmd/server/main.go` — entrypoint server (thin; boot logic di paket
    `app` donisgo).
  - `cmd/console/main.go` — entrypoint CLI console (paket `console` donisgo).
- `go.mod` — module `github.com/rolldone/donisgo-template`,
  require `github.com/rolldone/donisgo v0.1.4` (pinned, tanpa `replace`).
  `go.sum` ikut dibuat. Bump ke `v0.2.0` dilakukan di fase 5.

### Notes

- Import ke paket yang **tetap** di donisgo (`app`, `console`, kontrak
  `plugins`) tidak diubah — hanya isi aset yang dipindahkan.
- Registrasi plugin `pages` dari sisi donisgo (`internal/pluginloader`)
  masih berjalan di repo donisgo; pembersihan dilakukan pada fase 4.
