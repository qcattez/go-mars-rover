package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rover", func() {
	It("should create a rover with coordinates 0,0 and a north orientation", func() {
		rover, err := New(0, 0, "N")
		Expect(rover.Position()).To(Equal(Position{0, 0, "N"}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return an error when the orientation is empty", func() {
		_, err := New(0, 0, "")
		Expect(err).To(MatchError(InvalidOrientationError{}))
		Expect(err.Error()).To(Equal("The rover orientation is invalid"))
	})

	It("should return an error when the orientation is invalid", func() {
		_, err := New(0, 0, "A")
		Expect(err).To(MatchError(InvalidOrientationError{"A"}))
		Expect(err.Error()).To(Equal("\"A\" is not a valid rover orientation"))
	})

	It("should create a rover with coordinates 1,-2 and a north orientation", func() {
		rover, err := New(1, -2, "N")
		Expect(rover.Position()).To(Equal(Position{1, -2, "N"}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and a west orientation", func() {
		rover, err := New(0, 0, "W")
		Expect(rover.Position()).To(Equal(Position{0, 0, "W"}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and an east orientation", func() {
		rover, err := New(0, 0, "E")
		Expect(rover.Position()).To(Equal(Position{0, 0, "E"}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and a south orientation", func() {
		rover, err := New(0, 0, "S")
		Expect(rover.Position()).To(Equal(Position{0, 0, "S"}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should not move when there are no commands", func() {
		// Given
		rover, _ := New(0, 0, "N")
		commands := []Movement{}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, "N"}))
	})

	It("should move forward once with 1 forward command", func() {
		// Given
		rover, _ := New(0, 0, "N")
		commands := []Movement{Forward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, "N"}))
	})

	It("should move forward twice with 2 forward command", func() {
		// Given
		rover, _ := New(0, 0, "N")
		commands := []Movement{Forward, Forward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2, "N"}))
	})
})
