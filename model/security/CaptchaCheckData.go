package security

// CaptchaCheckData 结构体
type CaptchaCheckData struct {
	// 推荐的验证类型,0-放行 1-短信下行 2-语音验证 3-滑动验证 4-实人认证  32-综合（滑动＋语音） 9-阻断
	CaptchaType int64 `json:"captcha_type,omitempty" xml:"captcha_type,omitempty"`
	// 发起端上验证需要的信息
	CaptchaClientNeedInfo string `json:"captcha_client_need_info,omitempty" xml:"captcha_client_need_info,omitempty"`
}
