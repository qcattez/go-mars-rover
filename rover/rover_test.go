package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go-mars-rover/rover/movements"
	or "go-mars-rover/rover/orientations"
)

var _ = Describe("Rover", func() {
	It("should create a rover with coordinates 0,0 and a north orientation", func() {
		rover, err := New(0, 0, or.North)
		Expect(rover.Position()).To(Equal(Position{0, 0, or.North}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return an error when the orientation is empty", func() {
		_, err := New(0, 0, or.Unknown)
		Expect(err).To(MatchError(or.InvalidOrientationError{}))
		Expect(err.Error()).To(Equal("The rover orientation is invalid"))
	})

	It("should create a rover with coordinates 1,-2 and a north orientation", func() {
		rover, err := New(1, -2, or.North)
		Expect(rover.Position()).To(Equal(Position{1, -2, or.North}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and a west orientation", func() {
		rover, err := New(0, 0, or.West)
		Expect(rover.Position()).To(Equal(Position{0, 0, or.West}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and an east orientation", func() {
		rover, err := New(0, 0, or.East)
		Expect(rover.Position()).To(Equal(Position{0, 0, or.East}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a rover with coordinates 0,0 and a south orientation", func() {
		rover, err := New(0, 0, or.South)
		Expect(rover.Position()).To(Equal(Position{0, 0, or.South}))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should not move when there are no commands", func() {
		// Given
		rover, _ := New(0, 0, or.North)
		commands := []Movement{}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, or.North}))
	})

	It("should move forward once with 1 forward command", func() {
		// Given
		rover, _ := New(0, 0, or.North)
		commands := []Movement{Forward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, or.North}))
	})

	It("should move forward twice with 2 forward commands", func() {
		// Given
		rover, _ := New(0, 0, or.North)
		commands := []Movement{Forward, Forward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2, or.North}))
	})

	It("should move forward once with 2 forward and 1 backward commands", func() {
		// Given
		rover, _ := New(0, 0, or.North)
		commands := []Movement{Forward, Forward, Backward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, or.North}))
	})

	It("should move backward once with 1 backward command", func() {
		// Given
		rover, _ := New(0, 0, or.North)
		commands := []Movement{Backward}

		// When
		rover.Move(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, -1, or.North}))
	})
})
