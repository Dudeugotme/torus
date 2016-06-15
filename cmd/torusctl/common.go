package main

import (
	"fmt"
	"os"

	"github.com/coreos/torus"
	"github.com/coreos/torus/distributor"
	"github.com/coreos/torus/internal/flagconfig"
)

func die(why string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, why+"\n", args...)
	os.Exit(1)
}

func mustConnectToMDS() torus.MetadataService {
	cfg := torus.Config{
		MetadataAddress: etcdAddress,
	}
	mds, err := torus.CreateMetadataService("etcd", cfg)
	if err != nil {
		die("couldn't connect to etcd: %v", err)
	}
	return mds
}

func createServer() *torus.Server {
	cfg := flagconfig.BuildConfigFromFlags()
	cfg.MetadataAddress = etcdAddress
	srv, err := torus.NewServer(cfg, "etcd", "temp")
	if err != nil {
		die("Couldn't start: %s\n", err)
	}
	err = distributor.OpenReplication(srv)
	if err != nil {
		die("Couldn't start: %s", err)
	}
	return srv
}
