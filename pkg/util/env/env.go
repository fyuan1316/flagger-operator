package env

import (
	"fmt"
	"os"
)

var (
	defaultIstioNamespace = "istio-system"
)

const (
	ConfigIstioNamespace = "ASM_ISTIO_NAMESPACE"
)

func GetWatchNamespace() (string, error) {
	ns, found := os.LookupEnv("WATCH_NAMESPACE")
	if !found {
		return "", fmt.Errorf("WATCH_NAMESPACE must be set")
	}
	return ns, nil
}

func GetLeaderElectionNamespace() (string, bool) {
	return os.LookupEnv("LEADER_ELECTION_NAMESPACE")
}

func IstioNamespace() string {
	if ns := os.Getenv(ConfigIstioNamespace); ns != "" {
		return ns
	}
	return defaultIstioNamespace
}
