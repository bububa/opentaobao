package tbk

// TaobaotbksctpwdconvertMapData 结构体
type TaobaotbksctpwdconvertMapData struct {
	// 商品Id
	Numiid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品淘客转链链接
	Clickurl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 店铺卖家ID
	Sellerid string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 入参淘口令对应原始链接
	Originurl string `json:"origin_url,omitempty" xml:"origin_url,omitempty"`
	// 入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”
	Originpid string `json:"origin_pid,omitempty" xml:"origin_pid,omitempty"`
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景
	Bizsceneid string `json:"biz_scene_id,omitempty" xml:"biz_scene_id,omitempty"`
	// 短口令
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 长口令
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// 短链接
	Shorturl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 1-单品，2-店铺，3-活动，0-其它
	Urltype string `json:"url_type,omitempty" xml:"url_type,omitempty"`
}
