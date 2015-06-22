///Line 99, think about the structure of the entire file. 
//need getWildCard to take in the two arguments if you want it to be the main function
//I think you should chnage you main to point to an object instead of a function. An object with no attributes
//Go through each file in sample files. What is the relationship between main() and Run()?
//^My guess. Run is only for concurrency? When is Run called?
//Check out ui. 
//https://github.com/cloudfoundry/cli/blob/1c77293d3c4d5ae9f8374cfb173de35536b55f9b/cf/commands/application/app.go
//DRAW OUT after printing file ^, highlighting the variables that are important. 
//FIND THE LINK WITH ALL THE PLUGINS PEOPLE WROTE
//Q: The CLI will exit 0 if the plugin exits 0 and will exit
//*	1 should the plugin exits nonzero.
//https://github.com/cloudfoundry/cli/blob/master/plugin_examples/basic_plugin.go
//Q: Table only has "NewTable", "Add", "Print()"


package main

import (
	"errors"
	"fmt" //standard
	//"os"
	//"strconv"
	//"strings"
	//"time"

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

// type Printableable struct {
// 	ui            UI
// 	headers       []string
// 	headerPrinted bool
// 	maxSizes      []int
// 	rows          [][]string
// }
// func (t *Printableable) Add(row ...string) {
// 	t.rows = append(t.rows, row)
// }


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

	// if err := cmd.usage(args); nil != err { //usage is just confirmation for correct number of args
	// 	fmt.Println(err) //printing
	// 	os.Exit(1) //failuref
	// }
	output, _ := cliConnection.CliCommandWithoutTerminalOutput("apps")

	//table := terminal.NewTable(cmd.ui, []string{"",T("name"), T("requested state"), T("instances"), T("memory"), T("disk"), T("urls")})
	// //^ converts string to type T
	// for index, instance := range instances {
	// 	table.Add(
	// 		fmt.Spr
	// 		int("#%d", index),
	// 		fmt.Sprint("#%d", index),
	// 		fmt.Sprint("#%d", index),
	// 		fmt.Sprint("#%d", index),
	// 		fmt.Sprint("#%d", index),
	// 		fmt.Sprint("#%d", index),

	// 		)
	// }
	//table.Print()
	//output.Add("123", "3", "12", "23", "12", "12", "12")

	fmt.Println("first", len(output))
	fmt.Println("second", len(output[0]))
	fmt.Println(output[4])
	fmt.Println(output[5])
	fmt.Println("eeeeeee")
	for i := range output[4] {
		fmt.Println(output[i], ",")
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