package cloudgame

// GamepadGetResponse 结构体
type GamepadGetResponse struct {
	// 虚拟手柄样例
	VirtualGamepadList []VirtualGamepadList `json:"virtual_gamepad_list,omitempty" xml:"virtual_gamepad_list>virtual_gamepad_list,omitempty"`
}
