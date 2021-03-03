package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	l "github.com/swexbe/bagop/internal/pkg/logging"
	"github.com/swexbe/bagop/internal/pkg/utility"
)

func panicIfErr(err error) {
	if err != nil {
		l.Logger.Fatalf(err.Error())
	}
}

func main() {
	godotenv.Load()
	clean := flag.Bool("c", false, "Clean: Remove archives which have expired")
	backup := flag.Bool("b", false, "Backup: Make a backup and push it to Glacier")
	version := flag.Bool("version", false, "Version: Display version")
	verbose := flag.Bool("v", false, "Verbose: Display debug information")
	ttl := flag.String("ttl", "", "Time to Live: Number of days until archives will be deleted")
	vaultName := os.Getenv(utility.ENVVault)

	flag.Parse()

	if *verbose {
		l.Logger.SetLevel(logrus.DebugLevel)
		l.Logger.Infof("Running in verbose mode")
	}
	if *clean {
		cleanBackups(vaultName)
	} else if *backup {
		makeBackup(*ttl, vaultName)
	} else if *version {
		fmt.Printf("bagop v%s", utility.Version)
	} else {
		flag.PrintDefaults()
	}
}
