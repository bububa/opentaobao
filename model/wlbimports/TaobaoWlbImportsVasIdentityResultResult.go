package wlbimports

// TaobaowlbimportsvasidentityresultResult 结构体
type TaobaowlbimportsvasidentityresultResult struct {
	// 鉴定结果项
	VasResults []IdentityItemDto `json:"vas_results,omitempty" xml:"vas_results>identity_item_dto,omitempty"`
	// 物流订单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
}
