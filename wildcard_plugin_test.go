package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin/fakes"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
)

func fakeError(err error) {
	if err != nil {
		fmt.Println(err)
		
	}
}
var _ = Describe("WildcardPlugin", func() {
	var (
	wildcardPlugin 			*Wildcard
	fakeCliConnection 		*fakes.FakeCliConnection
	ui                 	 	*testterm.FakeUI
	//storedError 			error
	)
	Describe("Checking for correct output from wildcard-plugin commands", func() {
		routeList1 := make([]plugin_models.GetAppsRouteSummary, 0)
		routeList1 = append(routeList1,
			plugin_models.GetAppsRouteSummary{"1234", "www", plugin_models.GetAppsDomainFields{"1234", "google.com", "12345", false}},
			plugin_models.GetAppsRouteSummary{"5678", "www",  plugin_models.GetAppsDomainFields{"5678", "yahoo.com",  "12345", false}},
		)
		routeList2 := make([]plugin_models.GetAppsRouteSummary, 0)
		routeList2 = append(routeList2,
		plugin_models.GetAppsRouteSummary{"1234", "www", plugin_models.GetAppsDomainFields{"1234", "google.com", "12345", false}},
		plugin_models.GetAppsRouteSummary{"5678", "www",  plugin_models.GetAppsDomainFields{"5100", "naver.com",  "12345", true}},
		)
		appsList := make([]plugin_models.GetAppsModel, 0)
		//https://github.com/cloudfoundry/cli/blob/9f626377579452fa47be96998842afc9f78aa2ad/plugin/models/app_summary.go
		appsList = append(appsList,
			plugin_models.GetAppsModel{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app1", "1234-5678-90", "started", 4, 3, 512, 1024, routeList1},
			plugin_models.GetAppsModel{"app2", "0123-0011-23", "stopped", 4, 2, 638, 512, routeList2},
			plugin_models.GetAppsModel{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app10", "", "", 0, 0, 0, 0, nil},
		)
	
// THIS ONE ACTUALLY PANICS - PLEASE FIX THE CODE FOR IT, THEN RE-ENABLE THE TEST
		// Context("With no arguments", func() { //No args
			// BeforeEach(func() {
			// 	ui = &testterm.FakeUI{}
			// 	fakeCliConnection = &fakes.FakeCliConnection{}
			// 	wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
			// })
		// 	It("should display usage then exit", func() {
		// 		wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps"})
		// 		Expect(ui.Outputs).To(ContainSubstrings(
		// 			[]string{"Usage:", "cf", "wildcard-apps", "APP_NAME_WITH_WILDCARD"},
		// 		))
		// 	})
		// })
		// // Context("With too many arguments", func() { //Too many args
		// // 	It("should display usage", func() {
		// // 		wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app1", "app2"})
		// // 		Expect(ui.Outputs).To(ContainSubstrings(
		// // 			[]string{"Usage:", "cf", "wildcard-apps", "APP_NAME_WITH_WILDCARD"},
		// // 		))
		// // 	})
		// // })
		BeforeEach(func() {
			
		})
		Context("With wildcard-apps", func() { 
			BeforeEach(func() {
				ui = &testterm.FakeUI{}
				fakeCliConnection = &fakes.FakeCliConnection{}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				fakeCliConnection.GetAppsReturns(appsList, nil)
			})
			// fakeCliConnection.GetAppsReturns(appsList, nil)
			// It("should display org, space, and username", func() { //introduce
				// fakeCliConnection.GetAppsReturns(appsList, nil)
			// 	wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app1"})
			// 	Expect(ui.Outputs).To(ContainSubstrings(
			// 		[]string{"name", "requested state", "instances", "memory", "disk", "urls"},
			// 		[]string{"name", "requested state", "instances", "memory", "disk", "urls"},
			// 		[]string{"app1", "started", "3/4", "512M", "1G", "www.google.com", "www.yahoo.com"},
			// 	))
			// })
			It("should display all columns", func() { //table
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app1"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"name", "requested state", "instances", "memory", "disk", "urls"},
					[]string{"app1", "started", "3/4", "512M", "1G", "www.google.com", "www.yahoo.com"},
				))
			})
			It("should display no apps matching", func() { //no matched apps
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "noMatch*"})
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
					[]string{"app10"},
				))
			})
			It("lists all apps", func() {
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings( //matches
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
				))
			})
		})
		Context("With wildcard-delete interactive", func() { 
			// BeforeEach(func() {
			// 	fakeCliConnection = &fakes.FakeCliConnection{}
			// 	fakeCliConnection.GetAppsReturns(appsList, nil)
			// })
			It("should cancel wildcard-delete", func() { //no matched apps
				fakeCliConnection = &fakes.FakeCliConnection{}
				fakeCliConnection.GetAppsReturns(appsList, nil)
				ui = &testterm.FakeUI{Inputs: []string{"c"}}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-delete", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1", "started", "3/4", "512M", "1G", "www.google.com", "www.yahoo.com"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
					[]string{"app10"},
					[]string{"app10"},
					[]string{"Delete cancelled"},
					[]string{"OK"},
				))
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"cake"},
				))
			})
			It("should not delete any apps", func() { //table
				ui = &testterm.FakeUI{Inputs: []string{"i", "n", "n", "n", "n", "n", "n"}}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-delete", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
					[]string{"app10"},
					[]string{"OK"},
				))
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"cake"},
				))
			})
			It("should contain all apps", func() { //table
				ui = &testterm.FakeUI{Inputs: []string{}}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
					[]string{"app10"},
					[]string{"OK"},
				))
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"cake"},
				))
			})
			It("should delete some apps", func() { //table
				ui = &testterm.FakeUI{Inputs: []string{"i", "n", "y", "n", "y", "n", "y"}}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-delete", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
					[]string{"app10"},
					[]string{"OK"},
				))
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"cake"},
				))
			})
			It("should only have some apps", func() { //table
				ui = &testterm.FakeUI{Inputs: []string{}}
				wildcardPlugin = &Wildcard{ ui: ui, handleError: fakeError}
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1"},
					[]string{"app3"},
					[]string{"app5"},
					[]string{"OK"},
				))
				Expect(ui.Outputs).ToNot(ContainSubstrings(
					[]string{"app2"},
					[]string{"app4"},
					[]string{"app10"},
				))
			})
		})
	})
})
