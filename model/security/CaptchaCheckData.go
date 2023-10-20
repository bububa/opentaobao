package security

import (
	"sync"
)

// CaptchaCheckData 结构体
type CaptchaCheckData struct {
	// 发起端上验证需要的信息
	CaptchaClientNeedInfo string `json:"captcha_client_need_info,omitempty" xml:"captcha_client_need_info,omitempty"`
	// 推荐的验证类型,0-放行 1-短信下行 2-语音验证 3-滑动验证 4-实人认证  32-综合（滑动＋语音） 9-阻断
	CaptchaType int64 `json:"captcha_type,omitempty" xml:"captcha_type,omitempty"`
}

var poolCaptchaCheckData = sync.Pool{
	New: func() any {
		return new(CaptchaCheckData)
	},
}

// GetCaptchaCheckData() 从对象池中获取CaptchaCheckData
func GetCaptchaCheckData() *CaptchaCheckData {
	return poolCaptchaCheckData.Get().(*CaptchaCheckData)
}

// ReleaseCaptchaCheckData 释放CaptchaCheckData
func ReleaseCaptchaCheckData(v *CaptchaCheckData) {
	v.CaptchaClientNeedInfo = ""
	v.CaptchaType = 0
	poolCaptchaCheckData.Put(v)
}
