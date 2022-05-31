package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go-mars-rover/rover/commands"
	. "go-mars-rover/rover/orientations"
)

var _ = Describe("Rover", func() {
	It("should create a rover with coordinates 0,0 and a north orientation", func() {
		rover := New(0, 0, North)
		Expect(rover.Position()).To(Equal(Position{0, 0, North}))
	})

	It("should create a rover with coordinates 1,-2 and a west orientation", func() {
		rover := New(1, -2, West)
		Expect(rover.Position()).To(Equal(Position{1, -2, West}))
	})

	It("should not move when there are no commands", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, North}))
	})

	It("should move forward once with 1 forward command", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, North}))
	})

	It("should move forward twice with 2 forward commands", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{Forward, Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2, North}))
	})

	It("should move forward once with 2 forward and 1 backward commands", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{Forward, Forward, Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, North}))
	})

	It("should move backward once with 1 backward command", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, -1, North}))
	})

	It("should decrement the ordinate when moving forward oriented south", func() {
		// Given
		rover := New(0, 0, South)
		commands := []Command{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, -1, South}))
	})

	It("should increment the ordinate when moving backward oriented south", func() {
		// Given
		rover := New(0, 0, South)
		commands := []Command{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1, South}))
	})

	It("should increment the absiss when moving forward oriented west", func() {
		// Given
		rover := New(0, 0, West)
		commands := []Command{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 0, West}))
	})

	It("should decrement the absiss when moving backward oriented west", func() {
		// Given
		rover := New(0, 0, West)
		commands := []Command{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{-1, 0, West}))
	})

	It("should decrement the absiss when moving forward oriented east", func() {
		// Given
		rover := New(0, 0, East)
		commands := []Command{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{-1, 0, East}))
	})

	It("should increment the absiss when moving backward oriented east", func() {
		// Given
		rover := New(0, 0, East)
		commands := []Command{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 0, East}))
	})

	It("should be oriented east after a right turn from north orientation", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, East}))
	})

	It("should be oriented west after a left turn from north orientation", func() {
		// Given
		rover := New(0, 0, North)
		commands := []Command{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, West}))
	})

	It("should be oriented east after a left turn from south orientation", func() {
		// Given
		rover := New(0, 0, South)
		commands := []Command{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, East}))
	})

	It("should be oriented west after a right turn from south orientation", func() {
		// Given
		rover := New(0, 0, South)
		commands := []Command{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, West}))
	})

	It("should be oriented north after a left turn from east orientation", func() {
		// Given
		rover := New(0, 0, East)
		commands := []Command{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, North}))
	})

	It("should be oriented south after a right turn from east orientation", func() {
		// Given
		rover := New(0, 0, East)
		commands := []Command{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, South}))
	})

	It("should be oriented south after a left turn from west orientation", func() {
		// Given
		rover := New(0, 0, West)
		commands := []Command{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, South}))
	})

	It("should be oriented north after a right turn from west orientation", func() {
		// Given
		rover := New(0, 0, West)
		commands := []Command{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0, North}))
	})
})
