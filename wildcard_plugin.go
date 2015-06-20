package main

import (
	"errors"
	"fmt" //standard
	"os"
	//"strconv"
	//"strings"
	//"time"

	"github.com/cloudfoundry/cli/plugin" //standard

	//"path/filepath" //
	//"github.com/cloudfoundry/cli/cf/terminal"
)

//Wildcard is this plugin
type Wildcard struct {
	//matchedApps 	[]Apps
}
// func getMatchingApps() *Wildcard {
// 	targetsPath := filepath.Join(filepath.Dir(config_helpers.DefaultFilePath()), "targets")
// 	os.Mkdir(targetsPath, 0700)
// 	return &TargetsPlugin {
// 		configPath: config_helpers.DefaultFilePath(),
// 		targetsPath: targetsPath,
// 		currentPath: filepath.Join(targetsPath, "current"),
// 		suffix: "." + filepath.Base(config_helpers.DefaultFilePath()),
// 	}
// }

//GetMetadata returns metatada
func (cmd *Wildcard) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Wildcard",
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
// func newWildcard() *Wildcard {
// 	output, _ := cliConnection.CliCommandWithoutTerminalOutput("apps")
// 	table := terminal.NewTable(cmd.ui, []string{T("name"), T("requested state"), T("instances"), T("memory"), T("disk"), T("urls")})

// }

//Run runs the plugin
//called everytime user executes the comman
func (cmd *Wildcard) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "wildcard-apps" { //checking is very imp.
		cmd.WildcardCommandApps(cliConnection, args)
	} //else if args[0] == "wildcard-delete-a" {
	//	cmd.WildcardCommandDeleteAll(cliConnection, args)
	// } else if args[0] == "wildcard-delete-i" {
	// 	cmd.WildcardCommandDeleteInteractive(cliConnection, args)
	// }
}

//Q: what is the error for?
// func (cmd *Wildcard) getAppInformation(cliConnection plugin.CliConnection, name string) (AppStatus, error) {
// 	status := AppStatus {
// 		name:           	name,
// 		requested-state:   	0,
// 		instances: 			string,
// 		memory:          	string,
// 		disk:        		string,
// 		urls:				[]string{},
// 	}
// }

//WildcardCommand creates a new instance of this plugin
//this is the actual implementation
//one method per command
func (cmd *Wildcard) WildcardCommandApps(cliConnection plugin.CliConnection, args []string) {

	if err := cmd.usage(args); nil != err { //usage is just confirmation for correct number of args
		fmt.Println(err) //printing
		os.Exit(1) //failure
	}
	//returns 2 vals
	// rolloverTime, err := cmd.parseTime(args[3]) //go functions can ret multiple values
	// if nil != err {s
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// The getAppStatus calls will exit with an error if the named apps don't exist
	//lines 109 and 110 may be clearer than the ff
	// if cmd.app1, err = cmd.getAppStatus(cliConnection, args[1]); nil != err {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if cmd.app2, err = cmd.getAppStatus(cliConnection, args[2]); nil != err {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// cmd.showStatus()

	// count := cmd.app1.countRequested
	// if count == 0 {
	// 	fmt.Println("There are no instances of the source app to scale over")
	// 	os.Exit(0)  
	// }
	// sleepInterval := time.Duration(rolloverTime.Nanoseconds() / int64(count))

	// for count > 0 {
	// 	count--
	// 	cmd.app2.scaleUp(cliConnection) //
	// 	cmd.app1.scaleDown(cliConnection)
	// 	cmd.showStatus()
	// 	if count > 0 {
	// 		time.Sleep(sleepInterval)
	// 	}
	// }
	// fmt.Println()
}