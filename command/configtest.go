package command

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/agent"
)

// ConfigTestCommand is a Command implementation that is used to
// verify config files
type ConfigTestCommand struct {
	BaseCommand
}

func (c *ConfigTestCommand) Help() string {
	helpText := `
Usage: consul configtest [options]

  DEPRECATED. Use the 'consul validate' command instead.

  Performs a basic sanity test on Consul configuration files. For each file
  or directory given, the configtest command will attempt to parse the
  contents just as the "consul agent" command would, and catch any errors.
  This is useful to do a test of the configuration only, without actually
  starting the agent.

  Returns 0 if the configuration is valid, or 1 if there are problems.

` + c.BaseCommand.Help()

	return strings.TrimSpace(helpText)
}

func (c *ConfigTestCommand) Run(args []string) int {
	var configFiles []string

	f := c.BaseCommand.NewFlagSet(c)
	f.Var((*AppendSliceValue)(&configFiles), "config-file",
		"Path to a JSON file to read configuration from. This can be specified multiple times.")
	f.Var((*AppendSliceValue)(&configFiles), "config-dir",
		"Path to a directory to read configuration files from. This will read every file ending in "+
			".json as configuration in this directory in alphabetical order.")

	if err := c.BaseCommand.Parse(args); err != nil {
		return 1
	}

	if len(configFiles) <= 0 {
		c.UI.Error("Must specify config using -config-file or -config-dir")
		return 1
	}

	_, err := agent.ReadConfigPaths(configFiles)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Config validation failed: %v", err.Error()))
		return 1
	}
	return 0
}

func (c *ConfigTestCommand) Synopsis() string {
	return "DEPRECATED. Use the validate command instead"
}
