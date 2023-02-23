package alihouse

// StorePunishDto 结构体
type StorePunishDto struct {
	// dto
	PunishDtos []PunishDto `json:"punish_dtos,omitempty" xml:"punish_dtos>punish_dto,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
