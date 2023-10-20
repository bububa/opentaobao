package media

import (
	"sync"
)

// UserInfo 结构体
type UserInfo struct {
	// 用户订购的图片空间容量
	OrderSpace string `json:"order_space,omitempty" xml:"order_space,omitempty"`
	// 已使用的图片空间容量
	UsedSpace string `json:"used_space,omitempty" xml:"used_space,omitempty"`
	// 剩余的图片空间容量
	RemainingSpace string `json:"remaining_space,omitempty" xml:"remaining_space,omitempty"`
	// 用户的可用容量，即订购量与免费量之和
	AvailableSpace string `json:"available_space,omitempty" xml:"available_space,omitempty"`
	// 图片空间的免费容量
	FreeSpace string `json:"free_space,omitempty" xml:"free_space,omitempty"`
	// 图片空间的订购有效期
	OrderExpiryDate string `json:"order_expiry_date,omitempty" xml:"order_expiry_date,omitempty"`
	// 用户自定义的水印参数，通过&#34;|&#34;分割开，如果用户没有定义则为&#34;&#34;&lt;br/&gt;具体水印参数组合方法，用&#34;|&#34;分开，顺序按&#34;是否全局设置|水印文字|是否文字水印优先|透明度|字体|字体大小|字体是否加粗|字体是否斜体|字体是否加下划线|字体颜色|旋转角度|是否带阴影|水印位置|图片水印URL|reference水印相对位置&#34; reference取值有左上（1）/中间（3）/右下（2）,其中的null代表为空
	WaterMark string `json:"water_mark,omitempty" xml:"water_mark,omitempty"`
}

var poolUserInfo = sync.Pool{
	New: func() any {
		return new(UserInfo)
	},
}

// GetUserInfo() 从对象池中获取UserInfo
func GetUserInfo() *UserInfo {
	return poolUserInfo.Get().(*UserInfo)
}

// ReleaseUserInfo 释放UserInfo
func ReleaseUserInfo(v *UserInfo) {
	v.OrderSpace = ""
	v.UsedSpace = ""
	v.RemainingSpace = ""
	v.AvailableSpace = ""
	v.FreeSpace = ""
	v.OrderExpiryDate = ""
	v.WaterMark = ""
	poolUserInfo.Put(v)
}
