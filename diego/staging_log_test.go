package diego

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	"github.com/cloudfoundry-incubator/diego-acceptance-tests/helpers/assets"
)

var _ = Describe("An application being staged with Diego", func() {
	var appName string

	BeforeEach(func() {
		appName = generator.RandomName()
	})

	AfterEach(func() {
		Eventually(cf.Cf("delete", appName, "-f")).Should(Exit(0))
	})

	It("has its staging log streamed during a push", func() {
		Eventually(cf.Cf("push", appName, "-p", assets.NewAssets().Standalone, "--no-start", "-b", ZIP_NULL_BUILDPACK), CF_PUSH_TIMEOUT).Should(Exit(0))
		Eventually(cf.Cf("set-env", appName, DIEGO_STAGE_BETA, "true")).Should(Exit(0))

		start := cf.Cf("start", appName)

		Eventually(start, CF_PUSH_TIMEOUT).Should(Exit(0))
		Expect(start).Should(Say("Downloaded app package"))
		Expect(start).Should(Say(`Staging\.\.\.`))
	})
})
