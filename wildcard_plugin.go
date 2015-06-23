//helpful link for match: https://golang.org/src/path/filepath/match_test.go
package main

import (
	"errors"
	//"fmt" //standard
	"os"
	//"reflect" //used to see type of object
	"strconv"
	"github.com/guidowb/cf-go-client/panic" //panics 
	"strings"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin" //standard//https://github.com/cloudfoundry/cli/blob/8c310da376377c53f001d916708c056ce1558959/plugin/plugin.go

	"path/filepath" //for matches//https://golang.org/pkg/path/filepath/
	"github.com/cloudfoundry/cli/cf/terminal" //for table || https://github.com/cloudfoundry/cli/blob/4a108fd21d6633b250f6d9f46e870967cae96ac0/cf/terminal/table.go
	//. "github.com/cloudfoundry/cli/cf/i18n"
)

//Wildcard is this plugin
type Wildcard struct {
	ui 				terminal.UI
	matchedApps 	[]plugin_models.ApplicationSummary
}

//GetMetadata returns metatada
func (cmd *Wildcard) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "wildcard",
		Version: plugin.VersionType{ //leavealone
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{  //****** array of command structures
			{
				Name:     "wildcard-apps",
				Alias:	  "wc-a",
				HelpText: "List all apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-apps APP_NAME_WITH_WILDCARD",
				},
			}, 
			{
				Name:     "wildcard-delete-a",
				Alias:	  "wc-da",
				HelpText: "Delete all apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-delete-a APP_NAME_WITH_WILDCARD",
				},
			},
			{
				Name:     "wildcard-delete-i",
				Alias:	  "wc-di",
				HelpText: "Interactively delete apps in the target space matching the wildcard",
				UsageDetails: plugin.Usage{
					Usage: "cf wildcard-delete-i APP_NAME_WITH_WILDCARD",
				},
			},
		},
	}
}

func main() { 
	plugin.Start(new(Wildcard))
	//plugin.Start(newWildcard())
}

func (cmd *Wildcard) usage(args []string) error {
	badArgs := 3 != len(args)
	if badArgs {
		return errors.New("Usage: cf wildcard-apps\n\tcf wildcard-apps APP_NAME_WITH_WILDCARD")
	}
	return nil
}

//Run runs the plugin
//called everytime user executes the command
func (cmd *Wildcard) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "wildcard-apps" { //checking is very imp.
		cmd.WildcardCommandApps(cliConnection, args)
	} //else if args[0] == "wildcard-delete-a" {
	//	cmd.WildcardCommandDeleteAll(cliConnection, args)
	// } else if args[0] == "wildcard-delete-i" {
	// 	cmd.WildcardCommandDeleteInteractive(cliConnection, args)
	// }
}



//WildcardCommand creates a new instance of this plugin
//this is the actual implementation
//one method per command
func (cmd *Wildcard) WildcardCommandApps(cliConnection plugin.CliConnection, args []string) {
	defer panic.HandlePanics()
	pattern := args[1]
	// if err := cmd.usage(args); nil != err { //usage is just confirmation for correct number of args
	// 	fmt.Println(err) //printing
	// 	os.Exit(1) //failuref
	// }
	output, _ := cliConnection.GetApps()
	for i := 0; i < (len(output)); i++ {
		ok, _ := filepath.Match(pattern, output[i].Name)
		if ok {
			cmd.matchedApps = append(cmd.matchedApps, output[i])
		}
	}
	// for i := 0; i < (len(cmd.matchedApps)); i++ {
	// 	fmt.Println(cmd.matchedApps[i].Name)
	// }
	//fmt.Println(reflect.TypeOf(cmd.matchedApps))
	cmd.ui = terminal.NewUI(os.Stdin, terminal.NewTeePrinter())
	table := terminal.NewTable(cmd.ui, []string{("name"), ("requested state"), ("instances"), ("memory"), ("disk"), ("urls")})
	for _, app := range cmd.matchedApps {
		 var urls []string
		for _, route := range app.Routes {
			urls = append(urls, route.Host)
		}
		table.Add(
			app.Name,
			app.State,
			strconv.Itoa(app.RunningInstances),
			strconv.FormatInt(app.Memory, 2),
			strconv.FormatInt(app.DiskQuota, 2),
			strings.Join(urls, ", "),
		)
		// if cmd.pluginCall {
		// 	cmd.populatePluginModel(cmd.matchedApps)
		// }
	}
	table.Print()

	
}