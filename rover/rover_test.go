package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rover", func() {
	Describe("New", func() {
		It("should create a rover with coordinates and a north orientation", func() {
			_, err := New(0, 0, "N")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error when the orientation is invalid", func() {
			_, err := New(0, 0, "A")
			Expect(err).To(MatchError(InvalidOrientationError{"A"}))
		})

		It("should create a rover with coordinates and a west orientation", func() {
			_, err := New(0, 0, "W")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a rover with coordinates and an east orientation", func() {
			_, err := New(0, 0, "E")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a rover with coordinates and a south orientation", func() {
			_, err := New(0, 0, "S")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("InvalidOrientationError", func() {
		It("should print a default error message", func() {
			message := InvalidOrientationError{}.Error()
			Expect(message).To(Equal("The rover orientation is invalid"))
		})

		It("should print a more descriptive error message with the orientation", func() {
			message := InvalidOrientationError{"A"}.Error()
			Expect(message).To(Equal("\"A\" is not a valid rover orientation"))
		})
	})

	Describe("Move", func() {
		It("should not move when there are no commands", func() {
			// Given
			rover, _ := New(0, 0, "N")
			commands := []string{}

			// When
			rover.Move(commands)

			// Then
			Expect(rover.Position()).To(Equal(Position{0, 0, "N"}))
		})
	})
})
