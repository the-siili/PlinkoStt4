A deterministic plinko game with a frontend from godot and backend in go.

This was a fun learning project, where we wanted to know how for example 
casinos can make games where you can't easily cheat. Here the server 
runs the physics simulation in milliseconds and only sends the result
and the animation data to the client. After that the client can animate
the ball falling and show the result. Adding a database and registering
the results wouldn't be hard, but this was school project so we didn't
bother to go so far.

Made in collaboration with Alvar Oras
