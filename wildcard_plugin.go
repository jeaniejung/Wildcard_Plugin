package main

import (
	// "errors"
	"flag"
	"fmt"
	"os"
	"github.com/jeaniejung/Wildcard_Plugin/table"
	"strconv"
	// "github.com/guidowb/cf-go-client/panic"
	"strings"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin"
	"path/filepath"
	// "github.com/cloudfoundry/cli/cf/terminal"
	// "github.com/cloudfoundry/cli/cf/trace"
	// . "github.com/cloudfoundry/cli/cf/i18n"
	// "github.com/cloudfoundry/cli/cf/i18n/detection"
	// "github.com/cloudfoundry/cli/cf/configuration/core_config"
	// "github.com/cloudfoundry/cli/cf/configuration/config_helpers"
	"github.com/cloudfoundry/cli/cf/formatters"
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
				Alias:	  "wc-a",
				HelpText: "List all apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-apps APP_NAME_WITH_WILDCARD",
				},
			}, 
			{
				Name:     "wildcard-delete",
				Alias:	  "wc-d",
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

func newWildcard() (*Wildcard) {
	return &Wildcard{}
}

func (cmd *Wildcard) Run(cliConnection plugin.CliConnection, args []string) {
	wildcardFlagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	force := wildcardFlagSet.Bool("f", false, "forces deletion of all apps matching APP_NAME_WITH_WILDCARD")
	//routes := wildcardFlagSet.Bool("r", false, "delete routes asssociated with APP_NAME_WITH_WILDCARD")
	// Parse starting from [1] because the [0]th element is the
	// name of the command and 
	err := wildcardFlagSet.Parse(args[2:])
	checkError(err)

	if args[0] == "wildcard-apps" && len(args) == 2{
		cmd.WildcardCommandApps(cliConnection, args)
	} else if args[0] == "wildcard-delete" && len(args) >= 2 && len(args) <= 4{
		cmd.WildcardCommandDelete(cliConnection, args, force)
	} else {
		usage(args)
	}
}

func checkError(err error) {
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage(args []string) {
	if args[0] == "wildcard-apps" {
		fmt.Println("Usage: cf wildcard-apps\n\tcf wildcard-apps APP_NAME_WITH_WILDCARD")
	} else if args[0] == "wildcard-delete" {
		fmt.Println("Usage: cf wildcard-delete\n\tcf wildcard-delete APP_NAME_WITH_WILDCARD")
	}
}

// func (cmd *Wildcard) GetMatchedApps(cliConnection plugin.CliConnection, args []string) ([]plugin_models.GetAppsModel) {
// 	if err := cmd.usage(args); err != nil {
// 		checkError(err)
// 	}
// 	cmd.pattern = args[1]
// 	cmd.introduce(cliConnection, args)
// 	output, _ := cliConnection.GetApps()
// 	for i := 0; i < (len(output)); i++ {
// 		ok, _ := filepath.Match(cmd.pattern, output[i].Name)
// 		if ok {
// 			cmd.matchedApps = append(cmd.matchedApps, output[i])
// 		}
// 	}
// 	if len(cmd.matchedApps) <= 0 {
// 		//case *errors.ModelNotFoundError:
// 		cmd.ui.Warn("Apps matching %s do not exist.", cmd.pattern)
// 		cmd.handleError(errors.New(""))
// 	}
// 	return cmd.matchedApps
// }
func getMatchedApps(cliConnection plugin.CliConnection, args []string) ([]plugin_models.GetAppsModel) {
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
			strconv.Itoa(app.RunningInstances) + "/" + strconv.Itoa(app.TotalInstances),
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

func (cmd *Wildcard) WildcardCommandDelete(cliConnection plugin.CliConnection, args []string, force *bool) {
	output := getMatchedApps(cliConnection, args)
	if !*force && len(output) > 1{
		cmd.WildcardCommandApps(cliConnection, args)
	}
	for _, app := range output {
		if *force {
			cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
			fmt.Println("Deleting app", app.Name)
		} else {
			var confirmation string
			fmt.Printf("Really delete the app %s?> ", app.Name)
			fmt.Scanf("%s", &confirmation)
			if strings.EqualFold(confirmation,"y") || strings.EqualFold(confirmation,"yes") {
				cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
				fmt.Println("Deleting app", app.Name)
			}
		}
	}
	if len(output) == 0 {
		fmt.Println("No apps found matching", args[1])
	}
	fmt.Println("Ok")
}





// func (cmd *Wildcard) WildcardCommandDelete(cliConnection plugin.CliConnection, args []string) {
// 	cmd.WildcardCommandApps(cliConnection, args)
// 	response := cmd.ui.Ask("Would you like to delete the apps (i)nteractively, (a)ll, or (c)ancel this command?")
// 	if !strings.EqualFold(response,"a") && !strings.EqualFold(response,"all") && !strings.EqualFold(response,"i") && !strings.EqualFold(response,"interactively") {
// 		cmd.ui.Warn(T("Delete cancelled"))
// 		cmd.handleError(errors.New(""))
// 	} else {
// 		for _, app := range cmd.matchedApps {
// 			if strings.EqualFold(response,"i") || strings.EqualFold(response,"interactively"){
// 				cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name)
// 			} else if strings.EqualFold(response,"a") || strings.EqualFold(response,"all") {
// 				confirmation := cmd.ui.Confirm("Really delete all apps matching %q?", cmd.pattern)
// 				if !confirmation {
// 					cmd.ui.Warn(T("Delete all cancelled"))
// 					cmd.handleError(errors.New(""))
// 				} else {
// 					fmt.Println("Deleting all apps matching %q ", cmd.pattern)
// 					cliConnection.CliCommandWithoutTerminalOutput("delete", app.Name, "-f")
// 				}
// 			} else {
// 				return
// 			}
// 		}
// 		cmd.ui.Ok()
// 	}
// }
