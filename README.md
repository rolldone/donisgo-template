# donisgo-template

Starter aplikasi & tooling untuk ekosistem
[donisgo](https://github.com/rolldone/donisgo) — library framework Go
(gin + GORM + plugin architecture).

Repo ini **bukan library**: jadikan titik awal proyek aplikasi (clone/fork, atau
GitHub **Use this template**), lalu kembangkan dengan mengimpor
`github.com/rolldone/donisgo/...` yang di-pin di `go.mod`.

## Peta ekosistem

| Repo | Peran | Cara konsumsi |
|---|---|---|
| [donisgo](https://github.com/rolldone/donisgo) | **Pure library** — API inti (`app`, `plugins`, `db`, `auth`, `events`, `keydb`, `storage`, `mail`, `secrets`, `uuid`, `console`) | di-pin `go get github.com/rolldone/donisgo@vX.Y.Z` |
| **donisgo-template** (repo ini) | **Starter aplikasi + tooling** — binary server/CLI, scaffold `plugin new`, plugin konten `pages`, docs aplikasi | clone / fork / *Use this template* |

Dokumentasi API library (interface `plugins.Plugin`, bus `events`, `storage`,
`auth`, dst.) ada di README repo donisgo. README ini fokus pada hal aplikasi.

## Isi starter

- `cmd/server` — binary HTTP server (thin; boot logic di `app` donisgo).
- `cmd/console` — CLI: `migrate`, `seed` (dari `console` donisgo) + `plugin new`
  (paket `console/` lokal milik starter).
- `console/plugin.go` — generator `plugin new` (skeleton aplikasi).
- `plugins/pages` — plugin konten: serving SPA dari `./sub_app/webapp/dist`,
  redirect `/spa` → `/app`, proxy homepage via `HOMEPAGE_URL`, wildcard
  `NoRoute`.
- `sub_app/webapp` — skeleton frontend (Vite + React) untuk SPA.
- `.env.example` — contoh konfigurasi aplikasi.

## Quickstart

1. Siapkan env:

```bash
cp .env.example .env
# edit .env sesuai kebutuhan
```

2. Jalankan migrasi (core lalu plugin):

```bash
go run ./cmd/console migrate up
```

3. Jalankan server:

```bash
go run ./cmd/server
```

4. (Opsional) Generate plugin baru:

```bash
go run ./cmd/console plugin new --id my-plugin
```

Generator membuat migrasi awal dengan prefix timestamp gaya Laravel:
`YYYY_MM_DD_HHMMSS_name.up.sql` / `.down.sql`.

## Environment variables

Lihat `.env.example` sebagai titik awal. Ringkasan variabel penting:

**App / URL**
- `APP_NAME`, `APP_ENV` (`development|staging|production`), `PORT` (default 8080).
- `APP_URL`, `ADMIN_URL`, `FRONT_URL` — URL frontend/aplikasi.
- `HOMEPAGE_URL` — landing page terpisah yang di-proxy plugin `pages` untuk
  path tak dikenal (selain `/api`, `/app`, `/spa`); kosong → fallback
  `index.html` SPA.

**Database (GORM)**
- `DB_TYPE` (`postgres|mysql|mariadb`), `DB_HOST`, `DB_PORT`, `DB_NAME`,
  `DB_USER`, `DB_PASSWORD`, `DB_SSLMODE` (Postgres).
- Pooling: `DB_MAX_OPEN_CONNS`, `DB_MAX_IDLE_CONNS`, `DB_CONN_MAX_LIFETIME_SEC`.
- Jangan commit secret; di produksi pakai secret manager.

**Auth / JWT**
- `AUTH_JWT_SECRET` — secret kanonis (HS256). Ganti dengan nilai acak 32+ byte.
- `JWT_ACCESS_EXP_SECONDS` (mis. 900), `JWT_REFRESH_EXP_SECONDS` (mis. 1209600).
- Legacy `JWT_SECRET` / `JWT_ACCESS_SECRET` / `JWT_REFRESH_SECRET` masih dibaca
  sebagai fallback (deprecated) — utamakan `AUTH_JWT_SECRET`.
- Detail API `auth` (instans `auth.New(auth.Config{...})` vs wrapper
  package-level deprecated) ada di README donisgo.

**Redis / KeyDB** — `KEYDB_HOST`, `KEYDB_PORT`, `KEYDB_PASS`, `KEYDB_DB`.

**SMTP** — `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, `SMTP_FROM`,
`SMTP_USE_TLS`, `SMTP_STARTTLS`.

**Storage** — `STORAGE_DRIVER` (`local|s3`), `STORAGE_ROOT` (local),
`STORAGE_PUBLIC_URL` (URL publik untuk file), `S3_*` bila `s3`.

**CORS** — `CORS_ALLOWED_ORIGINS` (daftar origin dipisah koma).

**Trusted proxies (`ClientIP`)**
- `TRUSTED_PROXIES` — daftar IP/CIDR (dipisah koma) yang boleh mengeset
  `X-Forwarded-For` / `X-Real-IP`; gin memakainya untuk `ClientIP()`.
- **Kosong (default) = trust none**: header forwarded diabaikan, `ClientIP()`
  mengembalikan socket peer. Ini default aman.
- **Wajib diisi bila berjalan di belakang reverse proxy** (nginx/caddy/traefik/
  NPM); kalau tidak, `ClientIP()` melaporkan IP proxy, sehingga rate-limit
  berbasis IP dan audit log jadi tidak berguna.
- Entri tidak valid → dilog `[WARNING]` dan seluruh daftar jatuh ke trust-none;
  startup tidak pernah gagal karena variabel ini.

**Misc** — `LOG_LEVEL`, `DOCKER_HOST_IP`.

## Console commands

```bash
# Semua command
go run ./cmd/console --help

# Migrasi database
go run ./cmd/console migrate [make|up|down|down-all|list]

# Generator plugin (starter)
go run ./cmd/console plugin new --id <plugin-id> [--template minimal|crud|middleware]

# Seed data
go run ./cmd/console seed [--plugin core|all|<plugin-id>]
```

### Migrasi

File migrasi berada di `migrations/{db_type}` (core) dan
`plugins/{plugin_id}/migrations/{db_type}` (plugin), dengan `{db_type}` =
`postgres`, `mysql`, dst.

```bash
# buat pasangan migrasi up/down (timestamp otomatis)
go run ./cmd/console migrate make add_users_table

# terapkan (core dulu, lalu plugin)
go run ./cmd/console migrate up

# hanya untuk satu plugin
go run ./cmd/console migrate --plugin myplugin up

# rollback (plugin lebih dulu)
go run ./cmd/console migrate down
go run ./cmd/console migrate down-all

# status per target
go run ./cmd/console migrate list

# override tipe DB
go run ./cmd/console migrate --db mysql up
```

Catatan:
- Status migrasi disimpan di tabel `migrations` dan `migration_targets`
  (dibuat otomatis saat pertama jalan).
- Prefix timestamp UTC gaya Laravel menjaga urutan deterministik antar plugin
  dan environment.
- Target berstatus `dirty` → CLI menolak lanjut; periksa DB & file migrasi,
  lalu bersihkan `dirty` di `migration_targets`.
- Untuk produksi, tulis file SQL eksplisit dan uji rollback di staging dulu.

## Plugin development

### Generator (`plugin new`)

```bash
go run ./cmd/console plugin new --id my-plugin                    # minimal (default)
go run ./cmd/console plugin new --id my-plugin --template crud
go run ./cmd/console plugin new --id my-plugin --template middleware
go run ./cmd/console plugin new --id my-plugin --name "My Plugin"
go run ./cmd/console plugin new --id my-plugin --no-console
```

Template: `minimal` (handler health), `crud` (handler + service), `middleware`
(middleware contoh).

Struktur hasil:
- `plugins/my_plugin/plugin.go` — implementasi interface `plugins.Plugin`
- `plugins/my_plugin/handlers/` — handler HTTP
- `plugins/my_plugin/migrations/postgres/` — `<timestamp>_init.up.sql` / `.down.sql`
- `plugins/my_plugin/services/` — layer service (template `crud`)
- `plugins/my_plugin/middleware/` — implementasi middleware (template `middleware`)

Import framework di hasil scaffold **selalu** `github.com/rolldone/donisgo/...`
(pemisahan `frameworkModulePath`, fix #13), sedangkan paket milik aplikasi
(`handlers`, `middleware`, `services`) mengikuti module path `go.mod` di CWD.

### Registrasi plugin

Setelah generate, daftarkan plugin di kedua binary.

`cmd/server/main.go`:
```go
import (
   "github.com/rolldone/donisgo/app"
   "github.com/rolldone/donisgo/plugins"
   myplugin "github.com/rolldone/donisgo-template/plugins/my_plugin"
)

func main() {
   err := app.Run(app.Options{
      RegisterPlugins: func() {
         plugins.RegisterPlugins([]plugins.Plugin{myplugin.New()})
      },
   })
   // ...
}
```

`cmd/console/main.go`:
```go
console.RegisterAdditionalPlugins([]plugins.Plugin{myplugin.New()})
console.Execute()
```

Prefer wiring eksplisit `app.New(app.Config{Plugins: ...})` (lihat README
donisgo) di atas hook `app.Options` yang deprecated.

### Kontrak plugin

Implement `plugins.Plugin` (`plugins/types.go`):
- `ID() string`
- `RegisterServices(deps plugins.ServiceDeps) error` — `deps.DB`, `deps.Store`
- `RegisterMiddleware() []plugins.MiddlewareDescriptor` — target `global|admin|api`
- `RegisterRoutes(router *gin.Engine, admin, api *gin.RouterGroup) error`
- `Seed() error` — opsional
- `ConsoleCommands() []*cobra.Command` — opsional

### Plugin `pages` (serving halaman/SPA)

`plugins/pages` adalah plugin konten milik starter: melayani SPA dari
`./sub_app/webapp/dist` (hasil build Vite), redirect `/spa` → `/app`, proxy
homepage via `HOMEPAGE_URL`, dan menangani route tak dikenal lewat `NoRoute`.

- Karena ia memegang `router.NoRoute`, registrasinya harus **paling akhir**:
  gunakan `plugins.Registry.EnsurePluginLast` agar semua route spesifik
  (api/admin/health) menang lebih dulu. Idempotent — tidak dobel.
- **Sejak donisgo `v0.2.0` serving halaman tidak lagi otomatis**: aplikasi/sub-app
  yang mengandalkan `pages` WAJIB mendaftarkannya secara eksplisit dari
  `github.com/rolldone/donisgo-template/plugins/pages` (lihat CHANGELOG donisgo
  `v0.2.0`). Pada donisgo `v0.1.x` registrasi masih dilakukan framework.

### Storage: serving file lokal dari plugin

Bila ingin plugin menyajikan file statis dari storage lokal, daftarkan di
`RegisterRoutes`:

```go
if localStore, ok := p.service.Store.(*storage.LocalStore); ok {
   router.Static("/assets/products", localStore.GetRoot()+"/products")
}
```

Atur `STORAGE_PUBLIC_URL` agar tautan publik menunjuk ke lokasi yang benar
(local: `http://localhost:8080/assets`, S3/CDN: host publik bucket/CDN).

## DB & transaksi

- Koneksi GORM dibuka oleh `db.Open(db.ConfigFromEnv())`; inject handle ke
  service, jangan pakai variabel global.
- Bantu transaksi: `db.WithTransaction(ctx, gdb, fn)` (lihat `db/tx.go`).
- Teruskan `context.Context` dan, bila perlu, `*gorm.DB` transaksi dari handler
  ke service agar operasi ikut transaksi yang sama; gunakan `db.WithContext(ctx)`.
- Jaga transaksi tetap pendek; pastikan `Rollback()` dipanggil kecuali
  `Commit()` berhasil.
- `AutoMigrate` hanya untuk development; produksi pakai migrasi terstruktur.

## Security notes

- Jangan commit `.env` berisi secret; jaga `.env.example` tetap generik.
- Produksi: pakai secret manager (Vault, AWS Secrets Manager, dsb.) dan
  inject saat deploy.
- Rotasi key/secret berkala; prinsip least-privilege untuk akun DB/layanan.
- Isi `TRUSTED_PROXIES` bila di belakang reverse proxy.

## Menjaga sinkronisasi dengan donisgo

Repo ini mengonsumsi donisgo sebagai dependensi **pinned** (lihat `go.mod`).
Untuk naik versi: `go get github.com/rolldone/donisgo@vX.Y.Z && go mod tidy`,
lalu baca CHANGELOG donisgo (jalur `v0.x` membolehkan breaking pada bump
minor).

## Lisensi

MIT — lihat [LICENSE](LICENSE).
