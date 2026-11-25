// @title Donisgo API
// @version 1.0
// @description Admin API for donisgo - auto-generated swagger docs
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /admin
package main

import (
	"log"
	"syscall"

	pagesplugin "github.com/rolldone/donisgo-template/plugins/pages"
	"github.com/rolldone/donisgo/app"
	"github.com/rolldone/donisgo/plugins"
)

// main is a thin entrypoint; core boot logic lives in app.
//
// Sejak donisgo v0.2.0 (pure library) tidak ada lagi plugin yang
// didaftarkan otomatis — registrasi plugin 100% milik pemanggil via
// app.Config.Plugins. Plugin aplikasi didaftarkan pertama; plugin konten
// `pages` (SPA serving + wildcard catch-all) WAJIB didaftarkan PALING AKHIR —
// dijamin oleh plugins.Registry.EnsurePluginLast (catatan migrasi CHANGELOG
// donisgo v0.2.0).
//
// Field Config yang nil jatuh ke perilaku env seperti sebelumnya, jadi cukup
// isi yang ingin di-override.
func main() {
	// Set permissive umask so created files/dirs are group-writable by default
	syscall.Umask(0o002)

	application, err := app.New(app.Config{
		Plugins: []plugins.Plugin{
			// Register your app plugins here (before `pages`).
		},
	})
	if err != nil {
		log.Fatalf("failed to build app: %v", err)
	}

	// `pages` memegang router.NoRoute (SPA /app + wildcard) — selalu terakhir.
	application.GetRegistry().EnsurePluginLast(pagesplugin.New())

	log.Fatal(application.Run())
}
