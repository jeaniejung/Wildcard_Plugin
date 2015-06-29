//this file, wildcard_plugin_test.go, is created by '$ ginkgo generate wildcard_plugin.go'
//VOCAB:
//GetAppsStub = func() ([]plugin_models.ApplicationSummary, error)
//
package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin/fakes"
	//. "github.com/cloudfoundry/cli/testhelpers/matchers"
	//"github.com/onsi/gomega/matchers"
	// "fmt"
	// "reflect"
	//io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testapi "github.com/cloudfoundry/cli/cf/api/fakes"






	//. "github.com/cloudfoundry/cli/testhelpers/matchers"
)

//top-level describe container using Ginkgo's "Describe(text string, body func()) bool" cuntion.
//var_=.. allows us to eval Describe at the top level without the need to wrap it in "func init() {}"
var _ = Describe("WildcardPlugin", func() {

	Describe("Checking for correct results to getMatchedApps", func() {
		var (
			
			wildcardPlugin 		*Wildcard
			fakeCliConnection 	*fakes.FakeCliConnection
		)
		// runCommand := func(args ...string) bool {
		// 	wildcardPlugin.Run(fakeCliConnection, []string{"wc-a", "ca*"})
		// 	cmd := command_registry.Commands.FindCommand("apps")
		// 	return testcmd.RunCliCommand(cmd, args, requirementsFactory)
		// }
		appsList := make([]plugin_models.ApplicationSummary, 0)
		appsList = append(appsList,
			plugin_models.ApplicationSummary{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app1", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app10", "", "", 0, 0, 0, 0, nil},
			)

		BeforeEach(func() {
			
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = &Wildcard{}
		})

		Context("With wildcard asterisk(*)", func() {
			It("should return all apps starting with 'ca'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "ca*"})
				Expect(len(output)).To(Equal(3))
				Expect(output[0].Name).To(Equal("cake"))
				Expect(output[1].Name).To(Equal("carrot"))
				Expect(output[2].Name).To(Equal("car"))
			})
			It("should return all apps", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "*"})
				Expect(len(output)).To(Equal(15))
				Expect(output[0].Name).To(Equal("spring-music"))
				Expect(output[1].Name).To(Equal("spring-master"))
				Expect(output[2].Name).To(Equal("spring-nana"))
				Expect(output[3].Name).To(Equal("spring-spring"))
				Expect(output[4].Name).To(Equal("springtime"))
				Expect(output[5].Name).To(Equal("cake"))
				Expect(output[6].Name).To(Equal("carrot"))
				Expect(output[7].Name).To(Equal("car"))
				Expect(output[8].Name).To(Equal("c"))
				Expect(output[9].Name).To(Equal("app1"))
				Expect(output[10].Name).To(Equal("app2"))
				Expect(output[11].Name).To(Equal("app3"))
				Expect(output[12].Name).To(Equal("app4"))
				Expect(output[13].Name).To(Equal("app5"))
				Expect(output[14].Name).To(Equal("app10"))		
			})
			It("should return all apps starting with 'sp'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "sp*"})
				Expect(len(output)).To(Equal(5))
				Expect(output[0].Name).To(Equal("spring-music"))
				Expect(output[1].Name).To(Equal("spring-master"))
				Expect(output[2].Name).To(Equal("spring-nana"))
				Expect(output[3].Name).To(Equal("spring-spring"))
				Expect(output[4].Name).To(Equal("springtime"))				
			})
			It("should return all apps starting with 'app'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app*"})
				Expect(len(output)).To(Equal(6))
				Expect(output[0].Name).To(Equal("app1"))
				Expect(output[1].Name).To(Equal("app2"))
				Expect(output[2].Name).To(Equal("app3"))
				Expect(output[3].Name).To(Equal("app4"))
				Expect(output[4].Name).To(Equal("app5"))
				Expect(output[5].Name).To(Equal("app10"))
			})
		})
		Context("With wildcard question-mark(?)", func() {
			It("should return all apps with patter 'ca?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "ca?"})
				Expect(len(output)).To(Equal(1))
				Expect(output[0].Name).To(Equal("car"))
			})
			It("should return all apps with patter 'app?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app?"})
				Expect(len(output)).To(Equal(5))
				Expect(output[0].Name).To(Equal("app1"))
				Expect(output[1].Name).To(Equal("app2"))
				Expect(output[2].Name).To(Equal("app3"))
				Expect(output[3].Name).To(Equal("app4"))
				Expect(output[4].Name).To(Equal("app5"))
			})
			It("should return all apps with patter 'app?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app??"})
				Expect(len(output)).To(Equal(1))
				Expect(output[0].Name).To(Equal("app10"))
			})
		})
	})
	Describe("Checking for correct results to WildcardCommandApps", func() {
		appsList := make([]plugin_models.ApplicationSummary, 0)
		appsList = append(appsList,
			plugin_models.ApplicationSummary{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app1", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app10", "", "", 0, 0, 0, 0, nil},
		)
		var (
			wildcardPlugin *Wildcard
			fakeCliConnection *fakes.FakeCliConnection
			ui                  *testterm.FakeUI
			configRepo          core_config.Repository
			appSummaryRepo      *testapi.FakeAppSummaryRepo
			deps                command_registry.Dependency
		)
		runCommand := func(args ...string) bool {
			cmd := command_registry.Commands.FindCommand("apps")
			return testcmd.RunCliCommand(cmd, args, requirementsFactory)
		}
		updateCommandDependency := func(pluginCall bool) {
			deps.Ui = ui
			deps.Config = configRepo
			deps.RepoLocator = deps.RepoLocator.SetAppSummaryRepository(appSummaryRepo)
			command_registry.Commands.SetCommand(command_registry.Commands.FindCommand("apps").SetDependency(deps, pluginCall))
		}
		BeforeEach(func() {
			ui = &testterm.FakeUI{}
			deps.PluginModels.AppsSummary = &appsList
			updateCommandDependency(true)
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = &Wildcard{}
			appSummaryRepo = &testapi.FakeAppSummaryRepo{}

		})
		Context("With wildcard asterisk(*)", func() {
			It("should return all apps starting with 'ca'", func() {
				// appSummaryRepo.GetSummariesInCurrentSpaceApps = []models.Application{}
				runCommand()
				Ω(pluginAppModels[0].Name).To(Equal("Application-1"))






				// fakeCliConnection.GetAppsReturns(appsList, nil)
				// output := io_helpers.CaptureOutput(func() {
				// 	wildcardPlugin.Run(fakeCliConnection, []string{"wc-a", "ca*"})
				// })
				// fmt.Println(len(output))
				// fmt.Println(len(output[0]))
				// fmt.Println(output)
				// fmt.Println(ui.Outputs[1])
				// Expect(len(ui.Outputs[1])).To(Equal("10"))
				// Expect(ui.Outputs[1]).To(Equal("app10"))

				// Ω(ui.Outputs).Should(ConsistOf(
				// 	[]string{"Getting apps in", "my-org", "my-space", "my-user"},
				// 	[]string{"OK"},
				// 	[]string{"Application-1", "started", "?/2", "512M", "1G", "app1.cfapps.io"},
				// ))



				// output, _ := fakeCliConnection.CliCommandWithoutTerminalOutput("wc-a", "ca*")

				//fmt.Println(output)
				// fmt.Println(output[0])
				// fmt.Println(output[1])
				// fmt.Println(output[2])
				// fmt.Println(output[3])
				// fmt.Println(output[4])
				// fmt.Println(output[5])
				// fmt.Println(output[6])

				
				// fmt.Println(reflect.TypeOf(output))
				// fmt.Println(output[1][0])
				// fmt.Println(output[1][1])
				//  for idx, v := range output {
				// 	v = strings.TrimSpace(v)
				// 	if strings.HasPrefix(v, "FAILED") {
				// 		e := output[idx+1]
				// 		return status, errors.New(e)
				// 	}
				// 	if strings.HasPrefix(v, "requested state: ") {
				// 		status.state = strings.TrimPrefix(v, "requested state: ")
				// 	}
				// 	if strings.HasPrefix(v, "instances: ") {
				// 		instances := strings.TrimPrefix(v, "instances: ")
				// 		split := strings.Split(instances, "/")
				// 		status.countRunning, _ = strconv.Atoi(split[0])
				// 		status.countRequested, _ = strconv.Atoi(split[1])
				// 	}
				// 	if strings.HasPrefix(v, "urls: ") {
				// 		urls := strings.TrimPrefix(v, "urls: ")
				// 		status.routes = strings.Split(urls, ", ")
				// 	}
				// }

				// Expect(len(output)).To(Equal(3))
				// Expect(output[0].Name).To(Equal("cake"))
				// Expect(output[1].Name).To(Equal("carrot"))
				// Expect(output[2].Name).To(Equal("car"))


			})

			// It("should return all apps starting with 'sp'", func() {
			// 	fakeCliConnection.GetAppsReturns(appsList, nil)
			// 	var err error
			// 	_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
			// 	Expect(err.Error()).To(Equal("App app1 not found"))
			// })
		})
	// 	Context("With wildcard sp*", func() {
	// 		It("should return all apps starting with 'sp'", func() {
	// 			fakeCliConnection.GetAppsReturns(appsList, nil)
	// 			var err error
	// 			_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
	// 			Expect(err.Error()).To(Equal("App app1 not found"))
	// 		})
	// 		It("should return all apps starting with 'sp'", func() {
	// 			fakeCliConnection.GetAppsReturns(appsList, nil)
	// 			var err error
	// 			_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
	// 			Expect(err.Error()).To(Equal("App app1 not found"))
	// 		})
	// 	})
	// })
	// BeforeEach(func() {
	// 	fakeCliConnection = &fakes.FakeCliConnection{}
	// 	scaleoverCmdPlugin = &ScaleoverCmd{}
	})

})
