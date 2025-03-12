extends Label



func _physics_process(delta: float) -> void:
	global_position.y -= 60


func _on_timer_timeout() -> void:
	queue_free()
