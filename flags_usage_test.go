package flags_test

import (
	"github.com/simonleung8/flags"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Showing Flags Usage", func() {

	var (
		fc flags.FlagContext
	)

	BeforeEach(func() {
		fc = flags.New()
		fc.NewIntFlag("intFlag", "Usage for intFlag")
		fc.NewBoolFlag("boolFlag", "Usage for boolFlag")
		fc.NewBoolFlag("f", "Usage for f")
	})

	It("prefixes the flag name with spaces", func() {
		outputs := fc.ShowUsage(5)
		Ω(outputs).To(ContainSubstring("     -intFlag"))
		Ω(outputs).To(ContainSubstring("     -f"))
		Ω(outputs).To(ContainSubstring("     --boolFlag"))
	})

	It("prints the usages with non-bool flags first", func() {
		outputs := fc.ShowUsage(0)
		buffer := gbytes.BufferWithBytes([]byte(outputs))
		Eventually(buffer).Should(gbytes.Say("intFlag"))
		Eventually(buffer).Should(gbytes.Say("Usage for intFlag"))
		Eventually(buffer).Should(gbytes.Say("boolFlag"))
		Eventually(buffer).Should(gbytes.Say("Usage for boolFlag"))
		Ω(outputs).To(ContainSubstring("f"))
		Ω(outputs).To(ContainSubstring("Usage for f"))
	})

	It("prefixes the non-bool flag with '-'", func() {
		outputs := fc.ShowUsage(0)
		Ω(outputs).To(ContainSubstring("-intFlag"))
	})

	It("prefixes single character bool flags with '-'", func() {
		outputs := fc.ShowUsage(0)
		Ω(outputs).To(ContainSubstring("-f"))
	})

	It("prefixes multi-character bool flags with '--'", func() {
		outputs := fc.ShowUsage(0)
		Ω(outputs).To(ContainSubstring("--boolFlag"))
	})

	It("aligns the text by padding string with spaces", func() {
		outputs := fc.ShowUsage(0)
		Ω(outputs).To(ContainSubstring("-intFlag        Usage for intFlag"))
		Ω(outputs).To(ContainSubstring("-f              Usage for f"))
		Ω(outputs).To(ContainSubstring("--boolFlag      Usage for boolFlag"))
	})
})