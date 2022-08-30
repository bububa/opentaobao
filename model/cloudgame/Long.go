package cloudgame

// Long 结构体
type Long struct {
	// 玩家id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 位置索引
	PlayerIndex string `json:"player_index,omitempty" xml:"player_index,omitempty"`
}
