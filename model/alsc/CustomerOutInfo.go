package alsc

import (
	"sync"
)

// CustomerOutInfo 结构体
type CustomerOutInfo struct {
	// 外部类型  *      * 手机号           MOBILE(&#34;mobile&#34;,&#34;手机号注册&#34;),      *      * 微信openId           WECHAT(&#34;wechat&#34;,&#34;微信openid注册&#34;),      *      * 微信小程序Id           WEAPP(&#34;weapp&#34;,&#34;微信小程序注册&#34;),      *      * 支付宝用户ID           ALIPAY(&#34;alipay&#34;,&#34;支付宝id注册&#34;),      *      * 面部ID           FACE_CODE(&#34;faceCode&#34;,&#34;faceCode注册&#34;),      *      * 座机注册           PHONE_CUSTOMER(&#34;phone_customer&#34;,&#34;座机注册&#34;)
	OutType string `json:"out_type,omitempty" xml:"out_type,omitempty"`
	// 外部id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
}

var poolCustomerOutInfo = sync.Pool{
	New: func() any {
		return new(CustomerOutInfo)
	},
}

// GetCustomerOutInfo() 从对象池中获取CustomerOutInfo
func GetCustomerOutInfo() *CustomerOutInfo {
	return poolCustomerOutInfo.Get().(*CustomerOutInfo)
}

// ReleaseCustomerOutInfo 释放CustomerOutInfo
func ReleaseCustomerOutInfo(v *CustomerOutInfo) {
	v.OutType = ""
	v.OutId = ""
	poolCustomerOutInfo.Put(v)
}
