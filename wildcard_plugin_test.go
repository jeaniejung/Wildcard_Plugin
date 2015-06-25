//this file, wildcard_plugin_test.go, is created by '$ ginkgo generate wildcard_plugin.go'
//VOCAB:
//GetAppsStub = func() ([]plugin_models.ApplicationSummary, error)
//
package main

import (
	//. "github.com/jeaniejung/wildcard_plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin/fakes"
)

//top-level describe container using Ginkgo's "Describe(text string, body func()) bool" cuntion.
//var_=.. allows us to eval Describe at the top level without the need to wrap it in "func init() {}"
var _ = Describe("WildcardPlugin", func() {
	var (
		wildcardPlugin *Wildcard
		fakeCliConnection *fakes.FakeCliConnection
	)
	appsList := make([]plugin_models.ApplicationSummary, 10)

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		wildcardPlugin = &Wildcard{}
	})
	JustBeforeEach(func() {})

	Describe("Checking for correct results to wildcard-apps", func() {
		BeforeEach(func() {
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
		})
		Context("With wildcard sp*", func() {
			It("should return all apps starting with 'sp'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				var err error
				output, _  := wildcardPlugin.WildcardCommandApps(fakeCliConnection, []string{"wc-a", "sp*"})
				Expect(err.Error()).To(Equal("App app1 not found"))
			})
			// It("should return all apps starting with 'sp'", func() {
			// 	fakeCliConnection.GetAppsReturns(appsList, nil)
			// 	var err error
			// 	_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
			// 	Expect(err.Error()).To(Equal("App app1 not found"))
			// })
		})
		// Context("With wildcard sp*", func() {
		// 	It("should return all apps starting with 'sp'", func() {
		// 		fakeCliConnection.GetAppsReturns(appsList, nil)
		// 		var err error
		// 		_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
		// 		Expect(err.Error()).To(Equal("App app1 not found"))
		// 	})
		// 	It("should return all apps starting with 'sp'", func() {
		// 		fakeCliConnection.GetAppsReturns(appsList, nil)
		// 		var err error
		// 		_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
		// 		Expect(err.Error()).To(Equal("App app1 not found"))
		// 	})
		// })
	})
	// Describe("Checking for correct results to wildcard-apps", func() {
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
	// }

})
