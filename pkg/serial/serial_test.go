package serial

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Serial", func() {
	Describe("I22 Message Parsing", func(){
		It("Should parse an I22 message", func() {
			line := "!I22,1,023,2160,0,0,178,235,-,0,000e,1,0,023,2160,235,2,1,p,P"
			msg, err := parseZQI22(line)
			Expect(err).NotTo(HaveOccurred())
			Expect(msg.HDR).To(BeTrue())
			Expect(msg.SourceFrameRate).To(Equal(uint8(23)))
			Expect(msg.SourceAspectRatio).To(Equal(uint8(235)))
			Expect(msg.SourceVerticalResolution).To(Equal(uint16(2160)))
		})

		It("Should refuse to parse non-I22 messages", func() {
			line := "POWERON."
			_, err := parseZQI22(line)
			Expect(err).To(HaveOccurred())
		})
	})
})
