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
	// "image/color"
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

// func introduction(cliConnection plugin.CliConnection, args []string) {


// }
// func (cmd *Wildcard) introduce(cliConnection plugin.CliConnection, args []string) {
// 	currOrg, _ := cliConnection.GetCurrentOrg()
// 	currSpace, _ := cliConnection.GetCurrentSpace()
// 	currUsername, _ := cliConnection.Username()
// 	cmd.ui.Say(T("Getting apps in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...",
// 		map[string]interface{}{
// 			"OrgName":   terminal.EntityNameColor(currOrg.Name),
// 			"SpaceName": terminal.EntityNameColor(currSpace.Name),
// 			"Username":  terminal.EntityNameColor(currUsername)}))
// 	 cmd.ui.Ok()
// 	 cmd.ui.Say("")
// }

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
	// introduction()
	output := getMatchedApps(cliConnection, args)
	table := table.NewTable([]string{("name"), ("requested state"), ("instances"), ("memory"), ("disk"), ("urls")})
	for _, app := range output {
		var urls []string
		for _, route := range app.Routes {
			if route.Host == "" {
				urls = append(urls, route.Domain.Name)
			}
			urls = append(urls, fmt.Sprintf("%s.%s", route.Host, route.Domain.Name))
		}
		table.Add(
			app.Name,
			app.State,
			strconv.Itoa(app.RunningInstances)+"/"+strconv.Itoa(app.TotalInstances),
			formatters.ByteSize(app.Memory*formatters.MEGABYTE),
			formatters.ByteSize(app.DiskQuota*formatters.MEGABYTE),
			strings.Join(urls, ", "),
		)
	}
	table.Print()
	if len(output) == 0 {
		fmt.Println("No apps found matching", args[1])
	}
}

func (cmd *Wildcard) WildcardCommandDelete(cliConnection plugin.CliConnection, args []string, force *bool, routes *bool) {
	// introduction()
	output := getMatchedApps(cliConnection, args)
	if !*force && len(output) > 1 {
		cmd.WildcardCommandApps(cliConnection, args)
	}
	for _, app := range output {
		if *force && *routes {
			cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f", "-r")
			fmt.Println("Deleting app", app.Name, "and its mapped routes")
		} else if *force {
			cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
			fmt.Println("Deleting app", app.Name)
		} else {
			var confirmation string
			fmt.Printf("Really delete the app %s?> ", app.Name)
			fmt.Scanf("%s", &confirmation)
			if strings.EqualFold(confirmation, "y") || strings.EqualFold(confirmation, "yes") {
				if *routes {
					cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f", "-r")
					fmt.Println("Deleting app", app.Name, "and its mapped routes")
				} else {
					cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
					fmt.Println("Deleting app", app.Name)
				}
			}
		}
	}
	if len(output) == 0 {
		fmt.Println("No apps found matching", args[1])
	}
	fmt.Println("OK")
}
