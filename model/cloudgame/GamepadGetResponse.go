package cloudgame

import (
	"sync"
)

// GamepadGetResponse 结构体
type GamepadGetResponse struct {
	// 虚拟手柄样例
	VirtualGamepadList []VirtualGamepadList `json:"virtual_gamepad_list,omitempty" xml:"virtual_gamepad_list>virtual_gamepad_list,omitempty"`
}

var poolGamepadGetResponse = sync.Pool{
	New: func() any {
		return new(GamepadGetResponse)
	},
}

// GetGamepadGetResponse() 从对象池中获取GamepadGetResponse
func GetGamepadGetResponse() *GamepadGetResponse {
	return poolGamepadGetResponse.Get().(*GamepadGetResponse)
}

// ReleaseGamepadGetResponse 释放GamepadGetResponse
func ReleaseGamepadGetResponse(v *GamepadGetResponse) {
	v.VirtualGamepadList = v.VirtualGamepadList[:0]
	poolGamepadGetResponse.Put(v)
}
