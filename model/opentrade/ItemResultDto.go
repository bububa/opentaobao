package opentrade

// ItemResultDTO 
type ItemResultDTO struct {
    // 失败商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 失败原因
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
