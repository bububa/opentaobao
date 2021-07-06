package opentrade

// ItemResultDto 结构体
type ItemResultDto struct {
	// 失败原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 失败商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
