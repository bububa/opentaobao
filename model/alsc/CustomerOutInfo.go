package alsc

// CustomerOutInfo 结构体
type CustomerOutInfo struct {
	// 外部类型  *      * 手机号           MOBILE("mobile","手机号注册"),      *      * 微信openId           WECHAT("wechat","微信openid注册"),      *      * 微信小程序Id           WEAPP("weapp","微信小程序注册"),      *      * 支付宝用户ID           ALIPAY("alipay","支付宝id注册"),      *      * 面部ID           FACE_CODE("faceCode","faceCode注册"),      *      * 座机注册           PHONE_CUSTOMER("phone_customer","座机注册")
	OutType string `json:"out_type,omitempty" xml:"out_type,omitempty"`
	// 外部id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
}
