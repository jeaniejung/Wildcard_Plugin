package main

import (
	"flag"
	"fmt"
	"github.com/cloudfoundry/cli/cf/formatters"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/jeaniejung/Wildcard_Plugin/table"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Wildcard struct {
}

func (cmd *Wildcard) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "wildcard",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "wildcard-apps",
				Alias:    "wc-a",
				HelpText: "List all apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-apps APP_NAME_WITH_WILDCARD",
				},
			},
			{
				Name:     "wildcard-delete",
				Alias:    "wc-d",
				HelpText: "Delete apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-delete APP_NAME_WITH_WILDCARD",
				},
			},
		},
	}
}

func main() {
	plugin.Start(newWildcard())
}

func newWildcard() *Wildcard {
	return &Wildcard{}
}

func (cmd *Wildcard) Run(cliConnection plugin.CliConnection, args []string) {
	wildcardFlagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	force := wildcardFlagSet.Bool("f", false, "forces deletion of all apps matching APP_NAME_WITH_WILDCARD")
	routes := wildcardFlagSet.Bool("r", false, "delete routes asssociated with APP_NAME_WITH_WILDCARD")
	// Parse starting from [1] because the [0]th element is the
	// name of the command and
	err := wildcardFlagSet.Parse(args[2:])
	checkError(err)

	if args[0] == "wildcard-apps" && len(args) == 2 {
		cmd.WildcardCommandApps(cliConnection, args)
	} else if args[0] == "wildcard-delete" && len(args) >= 2 && len(args) <= 4 {
		if len(args) == 2 || *force || *routes {
			cmd.WildcardCommandDelete(cliConnection, args, force, routes)
		} else {
			usage(args)
		}
	} else {
		usage(args)
	}
}

func usage(args []string) {
	if args[0] == "wildcard-apps" {
		fmt.Println("Usage: cf wildcard-apps\n\tcf wildcard-apps APP_NAME_WITH_WILDCARD")
	} else if args[0] == "wildcard-delete" {
		fmt.Println("Usage: cf wildcard-delete\n\tcf wildcard-delete APP_NAME_WITH_WILDCARD")
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

func introduction(cliConnection plugin.CliConnection, args []string) {
	currOrg, _ := cliConnection.GetCurrentOrg()
	currSpace, _ := cliConnection.GetCurrentSpace()
	currUsername, _ := cliConnection.Username()
	fmt.Println("Getting apps matching", table.EntityNameColor(args[1]), "in org", table.EntityNameColor(currOrg.Name), "/ space", table.EntityNameColor(currSpace.Name), "as", table.EntityNameColor(currUsername))
	fmt.Println(table.SuccessColor("OK"))
	fmt.Println("")
}

func getMatchedApps(cliConnection plugin.CliConnection, args []string) []plugin_models.GetAppsModel {
	pattern := args[1]
	output, err := cliConnection.GetApps()
	checkError(err)
	matchedApps := []plugin_models.GetAppsModel{}
	for i := 0; i < (len(output)); i++ {
		ok, _ := filepath.Match(pattern, output[i].Name)
		if ok {
			matchedApps = append(matchedApps, output[i])
		}
	}
	return matchedApps
}

func (cmd *Wildcard) WildcardCommandApps(cliConnection plugin.CliConnection, args []string) {
	introduction(cliConnection, args)
	output := getMatchedApps(cliConnection, args)
	mytable := table.NewTable([]string{("name"), ("requested state"), ("instances"), ("memory"), ("disk"), ("urls")})
	for _, app := range output {
		var urls []string
		for _, route := range app.Routes {
			if route.Host == "" {
				urls = append(urls, route.Domain.Name)
			}
			urls = append(urls, fmt.Sprintf("%s.%s", route.Host, route.Domain.Name))
		}
		mytable.Add(
			app.Name,
			app.State,
			strconv.Itoa(app.RunningInstances)+"/"+strconv.Itoa(app.TotalInstances),
			formatters.ByteSize(app.Memory*formatters.MEGABYTE),
			formatters.ByteSize(app.DiskQuota*formatters.MEGABYTE),
			strings.Join(urls, ", "),
		)
	}
	mytable.Print()
	if len(output) == 0 {
		fmt.Println(table.WarningColor("No apps found matching"), table.WarningColor(args[1]))
	}
}

func (cmd *Wildcard) WildcardCommandDelete(cliConnection plugin.CliConnection, args []string, force *bool, routes *bool) {
	output := getMatchedApps(cliConnection, args)
	exit := false
	if !*force && len(output) > 1 {
		cmd.WildcardCommandApps(cliConnection, args)
		fmt.Println("")
		fmt.Printf("Would you like to delete the apps (%s)nteractively, (%s)ll, or (%s)ancel this command?%s", table.PromptColor("i"), table.PromptColor("a"), table.PromptColor("c"), table.PromptColor(">"))
		var mode string
		fmt.Scanf("%s", &mode)
		if strings.EqualFold(mode, "a") || strings.EqualFold(mode, "all") {
			*force = true
		} else if strings.EqualFold(mode, "i") || strings.EqualFold(mode, "interactively") {
		} else {
			fmt.Println(table.WarningColor("Delete cancelled"))
			exit = true
		}
	} else {
		introduction(cliConnection, args)
	}
	if !exit {
		for _, app := range output {
			coloredAppName := table.EntityNameColor(app.Name)
			if *force && *routes {
				cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f", "-r")
				fmt.Println("Deleting app", coloredAppName, "and its mapped routes")
			} else if *force {
				cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
				fmt.Println("Deleting app", coloredAppName)
			} else {
				var confirmation string
				fmt.Printf("Really delete the app %s?%s ", table.PromptColor(app.Name), table.PromptColor(">"))
				fmt.Scanf("%s", &confirmation)
				if strings.EqualFold(confirmation, "y") || strings.EqualFold(confirmation, "yes") {
					if *routes {
						cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f", "-r")
						fmt.Println("Deleting app", coloredAppName, "and its mapped routes")
					} else {
						cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
						fmt.Println("Deleting app", coloredAppName)
					}
				}
			}
		}
	}
	if len(output) == 0 {
		fmt.Println(table.WarningColor("No apps found matching"), table.WarningColor(args[1]))
	} else {
		fmt.Println(table.SuccessColor("OK"))
	}
}
