extends Area2D

@export var audio_player: AudioStreamPlayer

func _on_area_entered(area: Area2D) -> void:
	audio_player.play()
