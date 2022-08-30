package cloudgame

// TopQueryUserBodyDressRequest 结构体
type TopQueryUserBodyDressRequest struct {
	// 云游戏MixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// API版本号, 可选
	Version string `json:"version,omitempty" xml:"version,omitempty"`
}
