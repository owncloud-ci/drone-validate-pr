package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/drone/drone-go/plugin/validator"
	"github.com/owncloud-ci/drone-fork-approval/plugin"

	"github.com/sirupsen/logrus"
)

const (
	HTTPServerReadHeaderTimeout = 3 * time.Second
)

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

// spec provides the plugin settings.
type spec struct {
	Name         string
	Namespace    string
	BuildVersion string
	BuildDate    string
	Bind         string `env:"BIND" envDefault:":3000"`
	Secret       string `env:"SECRET,unset,required"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"Info"`
}

func main() {
	spec := &spec{
		Name:         filepath.Base(os.Args[0]),
		Namespace:    "DRONE_FORK_APPROVAL",
		BuildVersion: BuildVersion,
		BuildDate:    BuildDate,
	}

	version := flag.Bool("v", false, "prints version")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", spec.Name)
		flag.PrintDefaults()
		fmt.Println("\nConfiguration is read from environment variables:")
		fmt.Printf("  %s_LOG_LEVEL:\n", strings.ToUpper(spec.Namespace))
		fmt.Println("	log level")
		fmt.Println("	default: :Info")
		fmt.Printf("  %s_SECRET:\n", strings.ToUpper(spec.Namespace))
		fmt.Println("	shared secret which is used to authorize access")
		fmt.Println("	required")
		fmt.Printf("  %s_BIND:\n", strings.ToUpper(spec.Namespace))
		fmt.Println("	bind address the server is listening to")
		fmt.Println("	default: :3000")
	}

	flag.Parse()

	if *version {
		fmt.Printf("%v Version=%v BuildDate=%v\n", spec.Name, spec.BuildVersion, spec.BuildDate)
		os.Exit(0)
	}

	err := env.ParseWithOptions(spec, env.Options{
		Prefix: fmt.Sprintf("%s_", strings.ToUpper(spec.Namespace)),
	})
	if err != nil {
		logrus.Fatalf("failed to read configuration from %v", err.Error())
	}

	logLevel, logErr := logrus.ParseLevel(spec.LogLevel)
	if logErr != nil {
		logrus.Errorf("failed to set LogLevel %v", logErr.Error())

		logLevel = logrus.InfoLevel
		logrus.Infof("set log level %v as fallback", logLevel)
	} else {
		logrus.Infof("set log level %v", spec.LogLevel)
	}

	logrus.SetLevel(logLevel)

	handler := validator.Handler(
		spec.Secret,
		plugin.New(),
		logrus.StandardLogger(),
	)

	logrus.Infof("server listening on address %s", spec.Bind)

	pluginMux := http.NewServeMux()
	pluginMux.Handle("/", handler)

	server := &http.Server{
		Addr:              spec.Bind,
		Handler:           pluginMux,
		ReadHeaderTimeout: HTTPServerReadHeaderTimeout,
	}

	logrus.Fatal(server.ListenAndServe())
}
