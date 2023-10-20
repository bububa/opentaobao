package baodian

import (
	"sync"
)

// GameGiftRecordDto 结构体
type GameGiftRecordDto struct {
	// cp item id
	CpItemId string `json:"cp_item_id,omitempty" xml:"cp_item_id,omitempty"`
	// 记录id
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 记录状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolGameGiftRecordDto = sync.Pool{
	New: func() any {
		return new(GameGiftRecordDto)
	},
}

// GetGameGiftRecordDto() 从对象池中获取GameGiftRecordDto
func GetGameGiftRecordDto() *GameGiftRecordDto {
	return poolGameGiftRecordDto.Get().(*GameGiftRecordDto)
}

// ReleaseGameGiftRecordDto 释放GameGiftRecordDto
func ReleaseGameGiftRecordDto(v *GameGiftRecordDto) {
	v.CpItemId = ""
	v.RecordId = 0
	v.Status = 0
	poolGameGiftRecordDto.Put(v)
}
