package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go-mars-rover/rover/commands"
	. "go-mars-rover/rover/orientations"
)

var _ = Describe("Rover", func() {
	It("should create a rover with coordinates 0,0 and a north orientation", func() {
		rover := New(0, 0, 5, North)
		Expect(rover.Position()).To(Equal(Position{0, 0}))
		Expect(rover.Orientation()).To(Equal(North))
	})

	It("should create a rover with coordinates 1,-2 and a west orientation", func() {
		rover := New(1, 2, 5, West)
		Expect(rover.Position()).To(Equal(Position{1, 2}))
		Expect(rover.Orientation()).To(Equal(West))
	})

	It("should not move when there are no commands", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0}))
	})

	It("should move forward once with 1 forward command", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1}))
	})

	It("should move forward twice with 2 forward commands", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{Forward, Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2}))
	})

	It("should move forward once with 2 forward and 1 backward commands", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{Forward, Forward, Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1}))
	})

	It("should move backward once with 1 backward command", func() {
		// Given
		rover := New(0, 1, 5, North)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0}))
	})

	It("should decrement the ordinate when moving forward oriented south", func() {
		// Given
		rover := New(0, 1, 5, South)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0}))
	})

	It("should increment the ordinate when moving backward oriented south", func() {
		// Given
		rover := New(0, 0, 5, South)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 1}))
	})

	It("should increment the absiss when moving forward oriented west", func() {
		// Given
		rover := New(0, 0, 5, West)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 0}))
	})

	It("should decrement the absiss when moving backward oriented west", func() {
		// Given
		rover := New(1, 0, 5, West)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0}))
	})

	It("should decrement the absiss when moving forward oriented east", func() {
		// Given
		rover := New(1, 0, 5, East)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 0}))
	})

	It("should increment the absiss when moving backward oriented east", func() {
		// Given
		rover := New(0, 0, 5, East)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 0}))
	})

	It("should be oriented east after a right turn from north orientation", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(East))
	})

	It("should be oriented west after a left turn from north orientation", func() {
		// Given
		rover := New(0, 0, 5, North)
		commands := []CommandString{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(West))
	})

	It("should be oriented east after a left turn from south orientation", func() {
		// Given
		rover := New(0, 0, 5, South)
		commands := []CommandString{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(East))
	})

	It("should be oriented west after a right turn from south orientation", func() {
		// Given
		rover := New(0, 0, 5, South)
		commands := []CommandString{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(West))
	})

	It("should be oriented north after a left turn from east orientation", func() {
		// Given
		rover := New(0, 0, 5, East)
		commands := []CommandString{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(North))
	})

	It("should be oriented south after a right turn from east orientation", func() {
		// Given
		rover := New(0, 0, 5, East)
		commands := []CommandString{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(South))
	})

	It("should be oriented south after a left turn from west orientation", func() {
		// Given
		rover := New(0, 0, 5, West)
		commands := []CommandString{TurnLeft}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(South))
	})

	It("should be oriented north after a right turn from west orientation", func() {
		// Given
		rover := New(0, 0, 5, West)
		commands := []CommandString{TurnRight}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Orientation()).To(Equal(North))
	})

	It("should return at origin when moving forward to west at absiss 4 with planet circumference 5", func() {
		// Given
		rover := New(4, 2, 5, West)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2}))
	})

	It("should return at origin when moving forward to west at absiss 2 with planet circumference 3", func() {
		// Given
		rover := New(2, 2, 3, West)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{0, 2}))
	})

	It("should return at absiss 4 when moving backward from west at origin with planet circumference 5", func() {
		// Given
		rover := New(0, 2, 5, West)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{4, 2}))
	})

	It("should return at origin when moving forward to north at ordinate 4 with planet circumference 5", func() {
		// Given
		rover := New(1, 4, 5, North)
		commands := []CommandString{Forward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 0}))
	})

	It("should return at absiss 4 when moving backward from north at origin with planet circumference 5", func() {
		// Given
		rover := New(1, 0, 5, North)
		commands := []CommandString{Backward}

		// When
		rover.ExecuteCommands(commands)

		// Then
		Expect(rover.Position()).To(Equal(Position{1, 4}))
	})
})
