package alsc

// FcUrlRequest 结构体
type FcUrlRequest struct {
	// 链接
	ReplaceUrl string `json:"replace_url,omitempty" xml:"replace_url,omitempty"`
	// 用户id
	TbUserId string `json:"tb_user_id,omitempty" xml:"tb_user_id,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
}
