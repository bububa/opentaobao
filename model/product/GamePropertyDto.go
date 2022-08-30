package product

// GamePropertyDto 结构体
type GamePropertyDto struct {
	// 服务器id
	ServerId int64 `json:"server_id,omitempty" xml:"server_id,omitempty"`
	// 客户端id
	ClientId int64 `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 平台id
	PlatformId int64 `json:"platform_id,omitempty" xml:"platform_id,omitempty"`
	// 游戏id
	GameId int64 `json:"game_id,omitempty" xml:"game_id,omitempty"`
}
