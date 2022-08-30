package product

// GoodsServerDto 结构体
type GoodsServerDto struct {
	// 服务器名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 游戏ID
	GameId int64 `json:"game_id,omitempty" xml:"game_id,omitempty"`
	// 服务器ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
