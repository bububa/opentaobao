package tbk

// TaobaoTbkScTpwdConvertMapData 结构体
type TaobaoTbkScTpwdConvertMapData struct {
	// 商品Id
	NumIid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品淘客转链链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 店铺卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 入参淘口令对应原始链接
	OriginUrl string `json:"origin_url,omitempty" xml:"origin_url,omitempty"`
	// 入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”
	OriginPid string `json:"origin_pid,omitempty" xml:"origin_pid,omitempty"`
}
