# donisgo-template

Starter aplikasi & tooling untuk ekosistem
[donisgo](https://github.com/rolldone/donisgo) — library framework Go
(gin + GORM + plugin architecture).

Repo ini **bukan library**: gunakan sebagai titik awal proyek aplikasi
(clone/fork, atau GitHub **Use this template**), lalu kembangkan dengan
mengimpor `github.com/rolldone/donisgo/...` (di-pin via `go.mod`).

## Status

Repo lahir 2026-11-25 sebagai bagian dari
**SPEC-005 — donisgo Pure Library Split**: migrasi `cmd/server`,
`cmd/console`, `plugins/pages`, scaffold `plugin new`, dan dokumentasi
aplikasi keluar dari repo donisgo agar donisgo menjadi pure library.
Eksekusi per fase menyusul — rujukan spec di
[`donisgo/docs/specs`](https://github.com/rolldone/donisgo/tree/master/docs/specs).

## Rencana isi (pasca-SPEC-005)

- `cmd/server` — binary aplikasi starter
- `cmd/console` — CLI (migrate, seed, `plugin new`)
- `plugins/pages` — plugin konten (serving halaman)
- Skeleton `plugin new` — generator sub-app yang mengimpor donisgo
- Dokumentasi aplikasi: quickstart, env, deploy
