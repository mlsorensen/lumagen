package parsers

import (
	"github.com/mlsorensen/lumagen/pkg/serial/message"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parsers", func() {
	Describe("I22 Message Parsing", func(){

		It("Should parse an I22 message", func() {
			line := "!I22,1,023,2160,0,0,178,235,-,0,000e,1,0,023,2160,235,2,1,p,P"
			parser := ZQI22Parser{func(msg message.ZQI22Message) {
				Expect(msg.HDR).To(BeTrue())
				Expect(msg.SourceFrameRate).To(Equal(uint8(23)))
				Expect(msg.SourceAspectRatio).To(Equal(uint8(235)))
				Expect(msg.SourceVerticalResolution).To(Equal(uint16(2160)))
			}}
			err := parser.Parse(line)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should refuse to parse non-I22 messages", func() {
			line := "POWERON."
			parser := ZQI22Parser{nil}
			err := parser.Parse(line)
			Expect(err).To(HaveOccurred())
		})
	})
})
