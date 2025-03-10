package physics

import (
	"fmt"
	"math/rand/v2"

	"github.com/Alexander-r/box2d"
)

func GenerateGame() [][2]float64 {

	position_slice := [][2]float64{}

	fmt.Println(position_slice)

	// Define the gravity vector.
	gravity := box2d.MakeB2Vec2(0.0, -9.81)

	// Construct a world object, which will hold and simulate the rigid bodies.
	world := box2d.MakeB2World(gravity)

	generate_pin_map(&world)

	ball_body := generate_ball(&world)

	// Prepare for simulation. Typically we use a time step of 1/60 of a
	// second (60Hz) and 10 iterations. This provides a high quality simulation
	// in most game scenarios.
	timeStep := 1.0 / 60.0
	velocityIterations := 8
	positionIterations := 3
	// This is our little game loop.
	for {
		world.Step(timeStep, velocityIterations, positionIterations)

		pos := ball_body.GetPosition()

		position_slice = append(position_slice, [2]float64{pos.X, pos.Y})
		if ball_body.GetPosition().Y <= 0 {
			break
		}

	}

	return position_slice
}

func generate_ball(world *box2d.B2World) *box2d.B2Body {
	bd := box2d.MakeB2BodyDef()
	bd.Position.Set(-0.3+rand.Float64()*(0.3+0.3), 15.0)
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	bd.FixedRotation = true
	bd.AllowSleep = false

	ball_body := world.CreateBody(&bd)

	shape := box2d.MakeB2CircleShape()
	shape.M_radius = 0.25

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &shape
	fd.Density = 1
	fd.Friction = 0.2
	fd.Restitution = 0.6

	ball_body.CreateFixtureFromDef(&fd)
	return ball_body
}

func generate_pin_map(world *box2d.B2World) []*box2d.B2Body {

	pins := []*box2d.B2Body{}

	c := 10.0
	for y := 0.0; y < 14; y += 4 {
		for x := 1.0; x <= c; x += 2 {
			pins = append(pins, generate_pin(world, box2d.B2Vec2{X: x, Y: y}))
			pins = append(pins, generate_pin(world, box2d.B2Vec2{X: -x, Y: y}))

		}
		c -= 2
	}

	c = 5.0
	for y := 2.0; y <= 14; y += 4 {
		for x := 0.0; x <= c; x += 1 {
			pins = append(pins, generate_pin(world, box2d.B2Vec2{X: x * 2, Y: y}))

			if x != 0 {
				pins = append(pins, generate_pin(world, box2d.B2Vec2{X: -x * 2, Y: y}))
			}
		}
		c -= 1
	}

	return pins
}

func generate_pin(world *box2d.B2World, pos box2d.B2Vec2) *box2d.B2Body {
	bd := box2d.MakeB2BodyDef()
	bd.Position.Set(pos.X, pos.Y)
	bd.Type = box2d.B2BodyType.B2_staticBody
	bd.FixedRotation = true
	bd.AllowSleep = false

	pin_body := world.CreateBody(&bd)

	shape := box2d.MakeB2CircleShape()
	shape.M_radius = 0.25

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &shape
	fd.Density = 20.0
	pin_body.CreateFixtureFromDef(&fd)

	return pin_body
}
