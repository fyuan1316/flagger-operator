package env

import (
	"fmt"
	"os"
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
