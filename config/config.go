package config

import (
	"os"
)

var curConf *Config

// Config describes all values that can currently be configured for konf
type Config struct {
	KonfDir string
	Silent  bool
}

// This is mainly used to provide some sane and lively defaults for unit tests
// TODO with the new config system in place, it might make sense to rework all the test-cases and remove the "./konf" reference
func init() {
	curConf = &Config{
		KonfDir: "./konf",
	}
}

// ConfFromHomeDir returns an initialized config based on the users HomeDir
func ConfFromHomeDir() (*Config, error) {
	c := &Config{}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	c.KonfDir = home + "/.kube/konfs"
	c.Silent = false

	return c, nil
}

// InitWithOverrides sets the config to the config supplied as its argument
func InitWithOverrides(or *Config) {
	curConf = or
}

// Currently there is no need to customize store and active configs individually.
// Setting the konfDir should be enough

// ActiveDir returns the currently configured active directory
func ActiveDir() string {
	return curConf.KonfDir + "/active"
}

// StoreDir returns the currently configured store directory
func StoreDir() string {
	return curConf.KonfDir + "/store"
}

// LatestKonfFile returns the currently configured latest konf file
func LatestKonfFile() string {
	return curConf.KonfDir + "/latestkonf"
}
