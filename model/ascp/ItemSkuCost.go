package ascp

// ItemSkuCost 结构体
type ItemSkuCost struct {
	// skuId
	Skuid int64 `json:"skuid,omitempty" xml:"skuid,omitempty"`
	// itemId
	Itemid int64 `json:"itemid,omitempty" xml:"itemid,omitempty"`
	// 成本 分为单位
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 返回错误提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
