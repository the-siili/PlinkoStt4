extends Node

var http_scene: PackedScene = load("res://http_request.tscn")


func _on_button_pressed() -> void:
	$ButtonAudioPlayer.play()
	var new_request = http_scene.instantiate()
	get_tree().root.add_child(new_request)
	
