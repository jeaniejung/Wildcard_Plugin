package main

import (
	"errors"
	"fmt" //standard
	//"os"
	//"strconv"
	//"strings"
	//"time"
	//"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin" //standard//https://github.com/cloudfoundry/cli/blob/8c310da376377c53f001d916708c056ce1558959/plugin/plugin.go

	//"path/filepath" //
	//"github.com/cloudfoundry/cli/cf/terminal" //for table || https://github.com/cloudfoundry/cli/blob/4a108fd21d6633b250f6d9f46e870967cae96ac0/cf/terminal/table.go
	"github.com/cloudfoundry/cli/cf/api/app_instances"
)

//Wildcard is this plugin
type Wildcard struct {
	appInstancesRepo app_instances.AppInstancesRepository
	//matchedApps 	[]Apps
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
			// {
			// 	Name:     "wildcard-delete-a",
			// 	Alias:	  "wc-da",
			// 	HelpText: "Delete all apps in the target space matching the wildcard",
			// 	UsageDetails: plugin.Usage{
			// 		Usage: "cf wildcard-delete-a APP_NAME_WITH_WILDCARD",
			// 	},
			// },
			// {
			// 	Name:     "wildcard-delete-i",
			// 	Alias:	  "wc-di",
			// 	HelpText: "Interactively delete apps in the target space matching the wildcard",
			// 	UsageDetails: plugin.Usage{
			// 		Usage: "cf wildcard-delete-i APP_NAME_WITH_WILDCARD",
			// 	},
			// },
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

	// if err := cmd.usage(args); nil != err { //usage is just confirmation for correct number of args
	// 	fmt.Println(err) //printing
	// 	os.Exit(1) //failuref
	// }
	output, s := cliConnection.GetApps()
	fmt.Println(output)
	fmt.Println(output[0].Name)
	fmt.Println(output[1].Name)
	fmt.Println(s)
}