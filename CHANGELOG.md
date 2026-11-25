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

- **SPEC-005 fase 3** — terima scaffold + docs aplikasi dari repo `donisgo`
  (commit sumber `dfdd4f5`, 2026-11-25):
  - `console/` — paket console milik starter yang memuat generator
    `plugin new` (`console/plugin.go` + `console/plugin_test.go`, pindah utuh
    per OQ-3) dan meregistrasi command tree-nya ke root CLI framework
    (`github.com/rolldone/donisgo/console`).
  - `sub_app/webapp/` — skeleton frontend (Vite + React) yang dilayani plugin
    `pages` dari `./sub_app/webapp/dist`.
  - `.env.example` — konfigurasi aplikasi (App/DB/JWT/KeyDB/SMTP/Storage/CORS/
    TRUSTED_PROXIES), digabung & dirapikan dari `.env.example` donisgo.
  - README starter diperkaya dengan docs app-level: quickstart, env, console
    commands & migrasi, plugin development, plugin `pages`, storage, DB &
    transaksi, security notes.
- Import framework pada hasil scaffold `plugin new` selalu
  `github.com/rolldone/donisgo/...` (`frameworkModulePath`, fix #13); paket
  milik aplikasi mengikuti module path `go.mod`.

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

- Semua aset fase 3 dipindahkan **utuh tanpa refactor** (OQ-3);
  `migrate`/`seed` tidak diduplikasi (tetap memakai paket `console` donisgo).
- Pada fase 3 command `plugin` masih muncul **dua kali** di root CLI (milik
  donisgo `v0.1.x` + milik starter) — duplikasi sementara; implementasi
  donisgo dihapus pada fase 4.
- Import ke paket yang **tetap** di donisgo (`app`, `console`, kontrak
  `plugins`) tidak diubah — hanya isi aset yang dipindahkan.
- Registrasi plugin `pages` dari sisi donisgo (`internal/pluginloader`)
  masih berjalan di repo donisgo; pembersihan dilakukan pada fase 4.
