package interact

// InteractDto 结构体
type InteractDto struct {
	// 互动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 互动游戏app name
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 店铺ID
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 互动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 互动实例ID
	InteractId string `json:"interact_id,omitempty" xml:"interact_id,omitempty"`
	// 互动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 互动游戏app key
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 互动实例名称
	InstanceName string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
}
