package ascp

// EstConTimeModifyDto 结构体
type EstConTimeModifyDto struct {
	// 修改后的发货时效
	EstConTime string `json:"est_con_time,omitempty" xml:"est_con_time,omitempty"`
	// 修改前的发货时效
	OldEstConTime string `json:"old_est_con_time,omitempty" xml:"old_est_con_time,omitempty"`
	// 交易子单编号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 修改时间（时间戳）
	Modify int64 `json:"modify,omitempty" xml:"modify,omitempty"`
}
