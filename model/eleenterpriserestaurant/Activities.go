package eleenterpriserestaurant

// Activities 结构体
type Activities struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动文本描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// * 餐厅活动类型 	 * 3	优惠券（废弃） 	 * 7	新用户折扣（可以和其他活动共享) 	 * 100	额外折扣（废弃） 	 * 101	在线支付折扣（废弃） 	 * 102	活动折扣（类似满减的活动折扣） 	 * 103	新用户折扣（不可以和其他活动共享） 	 * 104	订单红包（废弃） 	 * 105	JINBAO折扣（专享红包，某些餐厅可以使用） 	 * 106	餐厅维度的赠品活动 	 * 美食活动类型 	 * 1	折扣价（按百分比折扣） 	 * 2	减价（按金额进行折扣） 	 * 3	第N份折扣价（第N件按百分比折扣） 	 * 4	固定价格（不管原价多少都按固定价格售卖，特价菜） 	 * 5	赠品（购物某个美食附送赠品，可以与其它美食活动并存）
	DetailType int64 `json:"detail_type,omitempty" xml:"detail_type,omitempty"`
	// 活动id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
