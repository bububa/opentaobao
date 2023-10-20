package ju

import (
	"sync"
)

// AlibabaJhsCommunitySubmittingTextData 结构体
type AlibabaJhsCommunitySubmittingTextData struct {
	// 微信直播二维码
	WechatLiveQR string `json:"wechat_live_q_r,omitempty" xml:"wechat_live_q_r,omitempty"`
	// 去微信直播
	ToWechatLive string `json:"to_wechat_live,omitempty" xml:"to_wechat_live,omitempty"`
	// 微信直播链接
	WechatLiveLink string `json:"wechat_live_link,omitempty" xml:"wechat_live_link,omitempty"`
	// 微信直播背景
	WechatLiveBackground string `json:"wechat_live_background,omitempty" xml:"wechat_live_background,omitempty"`
	// 去微信直播
	IsToWechatLive string `json:"is_to_wechat_live,omitempty" xml:"is_to_wechat_live,omitempty"`
	// 去淘宝背景
	ToTBBackground string `json:"to_t_b_background,omitempty" xml:"to_t_b_background,omitempty"`
	// 去淘宝头像
	ToTBAvatar string `json:"to_t_b_avatar,omitempty" xml:"to_t_b_avatar,omitempty"`
	// 二维码浮标
	QrFloat string `json:"qr_float,omitempty" xml:"qr_float,omitempty"`
	// 是否展示浮标
	IsShowQrFloat string `json:"is_show_qr_float,omitempty" xml:"is_show_qr_float,omitempty"`
	// 二维码图片
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 去淘宝
	ToTB string `json:"to_t_b,omitempty" xml:"to_t_b,omitempty"`
	// 淘口令
	NoTaoTokenTip string `json:"no_tao_token_tip,omitempty" xml:"no_tao_token_tip,omitempty"`
	// 提示语
	QrCodeTip string `json:"qr_code_tip,omitempty" xml:"qr_code_tip,omitempty"`
	// 提示语标题
	QrCodeTitle string `json:"qr_code_title,omitempty" xml:"qr_code_title,omitempty"`
	// 链接已复制
	OperationStatus string `json:"operation_status,omitempty" xml:"operation_status,omitempty"`
	// 就可以抢购啦
	OperationHint string `json:"operation_hint,omitempty" xml:"operation_hint,omitempty"`
}

var poolAlibabaJhsCommunitySubmittingTextData = sync.Pool{
	New: func() any {
		return new(AlibabaJhsCommunitySubmittingTextData)
	},
}

// GetAlibabaJhsCommunitySubmittingTextData() 从对象池中获取AlibabaJhsCommunitySubmittingTextData
func GetAlibabaJhsCommunitySubmittingTextData() *AlibabaJhsCommunitySubmittingTextData {
	return poolAlibabaJhsCommunitySubmittingTextData.Get().(*AlibabaJhsCommunitySubmittingTextData)
}

// ReleaseAlibabaJhsCommunitySubmittingTextData 释放AlibabaJhsCommunitySubmittingTextData
func ReleaseAlibabaJhsCommunitySubmittingTextData(v *AlibabaJhsCommunitySubmittingTextData) {
	v.WechatLiveQR = ""
	v.ToWechatLive = ""
	v.WechatLiveLink = ""
	v.WechatLiveBackground = ""
	v.IsToWechatLive = ""
	v.ToTBBackground = ""
	v.ToTBAvatar = ""
	v.QrFloat = ""
	v.IsShowQrFloat = ""
	v.QrCode = ""
	v.ToTB = ""
	v.NoTaoTokenTip = ""
	v.QrCodeTip = ""
	v.QrCodeTitle = ""
	v.OperationStatus = ""
	v.OperationHint = ""
	poolAlibabaJhsCommunitySubmittingTextData.Put(v)
}
