package diego

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
)

const (
	CF_PUSH_TIMEOUT                       = 4 * time.Minute
	LONG_CURL_TIMEOUT                     = 4 * time.Minute
	DOCKER_IMAGE_DOWNLOAD_DEFAULT_TIMEOUT = 5 * time.Minute

	ZIP_NULL_BUILDPACK = "https://github.com/cloudfoundry-incubator/null-buildpack/archive/master.zip"
	GIT_NULL_BUILDPACK = "https://github.com/cloudfoundry-incubator/null-buildpack"

	DIEGO_STAGE_BETA = "DIEGO_STAGE_BETA"
	DIEGO_RUN_BETA   = "DIEGO_RUN_BETA"
)

var context helpers.SuiteContext

func TestApplications(t *testing.T) {
	RegisterFailHandler(Fail)

	SetDefaultEventuallyTimeout(time.Minute)
	SetDefaultEventuallyPollingInterval(time.Second)

	config := helpers.LoadConfig()
	context = helpers.NewContext(config)
	environment := helpers.NewEnvironment(context)

	BeforeSuite(func() {
		environment.Setup()
	})

	AfterSuite(func() {
		environment.Teardown()
	})

	componentName := "Diego"

	rs := []Reporter{}

	if config.ArtifactsDirectory != "" {
		helpers.EnableCFTrace(config, componentName)
		rs = append(rs, helpers.NewJUnitReporter(config, componentName))
	}

	RunSpecsWithDefaultAndCustomReporters(t, componentName, rs)
}