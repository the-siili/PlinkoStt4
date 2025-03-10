extends HTTPRequest

var ball_scene: PackedScene = load("res://ball.tscn")

func _ready() -> void:
	var bet = float(get_tree().get_first_node_in_group("bet_box").get_line_edit().text)
	var payload = { "bet":bet }
	var json = JSON.stringify(payload)
	var headers = ["Content-Type: application/json"]
	request("http://localhost:8080/generate_game", headers, HTTPClient.METHOD_GET, json)

func _on_request_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var json = JSON.parse_string(body.get_string_from_utf8())
	var new_ball = ball_scene.instantiate()
	new_ball.position = Vector2(0, 1500.0)
	new_ball.positions = json["positions"]
	print(response_code)
	print(json["multiplier"])
	get_tree().root.add_child(new_ball)
