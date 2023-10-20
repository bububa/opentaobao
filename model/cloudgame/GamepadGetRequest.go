package cloudgame

import (
	"sync"
)

// GamepadGetRequest 结构体
type GamepadGetRequest struct {
	// xxxx
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
}

var poolGamepadGetRequest = sync.Pool{
	New: func() any {
		return new(GamepadGetRequest)
	},
}

// GetGamepadGetRequest() 从对象池中获取GamepadGetRequest
func GetGamepadGetRequest() *GamepadGetRequest {
	return poolGamepadGetRequest.Get().(*GamepadGetRequest)
}

// ReleaseGamepadGetRequest 释放GamepadGetRequest
func ReleaseGamepadGetRequest(v *GamepadGetRequest) {
	v.MixGameId = ""
	poolGamepadGetRequest.Put(v)
}
