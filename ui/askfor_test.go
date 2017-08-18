package ui_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Askfor", func() {
	Context("when using a console UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
		})
		It("AskForText should return the text value", func() {
			bIn.Write([]byte("testing\n"))
			s := ui.AskForText("Please type test:")
			Ω(bOut).To(gbytes.Say("Please type test:"))
			Ω(s).Should(Equal("testing"))
		})
		It("AskForInt should return the int value", func() {
			bIn.Write([]byte("12\n"))
			i, _ := ui.AskForInt("Please type an int:")
			Ω(bOut).To(gbytes.Say("Please type an int:"))
			Ω(i).Should(Equal(12))
		})
		// FIt("AskForInt should display an error if the value is not an int", func() {
		// 	ui.SetMaxInvalidInput(1)
		// 	bIn.Write([]byte("a\n"))
		// 	i, err := ui.AskForInt("Please type an int:")
		// 	Ω(bOut).To(gbytes.Say("Please type an int:"))
		// 	Eventually(bOut).Should(gbytes.Say("invalid"))
		// 	log.Printf("value %v", i)
		// 	//Ω(i).Should(Equal(0))
		//
		// })
		It("AskForFloat should return the float value", func() {
			bIn.Write([]byte("120.0\n"))
			f := ui.AskForFloat("Please type a float:")
			Ω(bOut).To(gbytes.Say("Please type a float:"))
			Ω(f).Should(Equal(120.0))
		})

	})

})
