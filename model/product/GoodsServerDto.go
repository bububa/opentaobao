package product

import (
	"sync"
)

// GoodsServerDto 结构体
type GoodsServerDto struct {
	// 服务器名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 游戏ID
	GameId int64 `json:"game_id,omitempty" xml:"game_id,omitempty"`
	// 服务器ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolGoodsServerDto = sync.Pool{
	New: func() any {
		return new(GoodsServerDto)
	},
}

// GetGoodsServerDto() 从对象池中获取GoodsServerDto
func GetGoodsServerDto() *GoodsServerDto {
	return poolGoodsServerDto.Get().(*GoodsServerDto)
}

// ReleaseGoodsServerDto 释放GoodsServerDto
func ReleaseGoodsServerDto(v *GoodsServerDto) {
	v.Name = ""
	v.GameId = 0
	v.Id = 0
	poolGoodsServerDto.Put(v)
}
