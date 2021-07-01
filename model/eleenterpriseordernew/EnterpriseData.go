package eleenterpriseordernew

// EnterpriseData 结构体
type EnterpriseData struct {
	// 预定送达时间
	DeliverTime string `json:"deliver_time,omitempty" xml:"deliver_time,omitempty"`
	// 送餐地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 收餐人姓名
	Consignee string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 菜价加上配送费和打包费的价格
	OriginalPrice int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 饿了么订单Id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 订单备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 电话号码
	PhoneList []string `json:"phone_list,omitempty" xml:"phone_list>string,omitempty"`
	// 餐厅唯一码
	OnlyRestaurantCode string `json:"only_restaurant_code,omitempty" xml:"only_restaurant_code,omitempty"`
	// 订单创建时间
	CreatedAt string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 餐厅名称
	RestaurantName string `json:"restaurant_name,omitempty" xml:"restaurant_name,omitempty"`
	// 长餐厅Id
	ErestaurantId string `json:"erestaurant_id,omitempty" xml:"erestaurant_id,omitempty"`
	// 送达费用
	DeliverFee string `json:"deliver_fee,omitempty" xml:"deliver_fee,omitempty"`
	// 订单状态码
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 配送员信息
	DeliverymanInfo *DeliverymanInfoDto `json:"deliveryman_info,omitempty" xml:"deliveryman_info,omitempty"`
	// 状态
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 最近更新时间
	LastUpdatedAt string `json:"last_updated_at,omitempty" xml:"last_updated_at,omitempty"`
	// 地址信息
	TrackingInfo *TrackingInfoDto `json:"tracking_info,omitempty" xml:"tracking_info,omitempty"`
}
