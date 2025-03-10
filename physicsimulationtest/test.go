package main

import (
	"fmt"

	"github.com/Alexander-r/box2d"
)

type Position struct {
	x box2d.B2Vec2
	y box2d.B2Vec2
}

func main() {

	position_slice := []Position{}

	fmt.Println(position_slice)

	// Define the gravity vector.
	gravity := box2d.MakeB2Vec2(0.0, -9.81)

	// Construct a world object, which will hold and simulate the rigid bodies.
	world := box2d.MakeB2World(gravity)

	bd := box2d.MakeB2BodyDef()
	bd.Position.Set(3.0, 5.0)
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	bd.FixedRotation = true
	bd.AllowSleep = false

	ball_body := world.CreateBody(&bd)

	shape := box2d.MakeB2CircleShape()
	shape.M_radius = 0.5

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &shape
	fd.Density = 20.0
	ball_body.CreateFixtureFromDef(&fd)

	// Prepare for simulation. Typically we use a time step of 1/60 of a
	// second (60Hz) and 10 iterations. This provides a high quality simulation
	// in most game scenarios.
	timeStep := 1.0 / 60.0
	velocityIterations := 8
	positionIterations := 3

	// This is our little game loop.
	for {
		world.Step(timeStep, velocityIterations, positionIterations)

		//fmt.Println(ball_body.GetPosition())

		position_slice = append(position_slice, ball_body.GetPosition())
		if ball_body.GetPosition().Y <= 0 {
			break
		}

	}

	fmt.Println(position_slice)
}
