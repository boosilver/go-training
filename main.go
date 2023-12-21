package main

import (
	"os"

	"github.com/boosilver/go-training/config"
	"github.com/boosilver/go-training/modules/servers"
	"github.com/boosilver/go-training/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
