package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
)

type AgentConfig struct {
	WorkDir            string `json:"workDirectory,omitempty"`
	LocalPlanDir       string `json:"localPlanDirectory,omitempty"`
	RemoteEnabled      bool   `json:"remoteEnabled,omitempty"`
	ConnectionInfoFile string `json:"connectionInfoFile,omitempty"`
}

type ConnectionInfo struct {
	KubeConfig string `json:"kubeConfig"`
	Namespace  string `json:"namespace"`
	SecretName string `json:"secretName"`
}


func Parse(path string, result interface{}) error {
	if path == "" {
		return fmt.Errorf("empty file passed")
	}

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	file := filepath.Base(path)
	switch {
	case strings.Contains(file, ".json"):
		return json.NewDecoder(f).Decode(result)
	case strings.Contains(file, ".yaml"):
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(b, result)
	default:
		return fmt.Errorf("file %s was not a JSON or YAML file", file)
	}
	return nil
}

// Use the token for basic auth, post the config file....
// Send the information in JSON encoded file in a header
// POST the token and the other information, get back a connection info file
