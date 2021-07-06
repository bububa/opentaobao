package jst

// SmsTaskModel 结构体
type SmsTaskModel struct {
	// 短信文案
	Contents []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
	// 短信签名
	SignNames []string `json:"sign_names,omitempty" xml:"sign_names>string,omitempty"`
	// 短信模板code
	TemplateCodes []string `json:"template_codes,omitempty" xml:"template_codes>string,omitempty"`
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// isv淘宝nick
	IsvNick string `json:"isv_nick,omitempty" xml:"isv_nick,omitempty"`
	// 商家淘宝nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 任务对应的短信类型 ：1--数字短信  2--权益短信  3--公众号短信
	SmsType string `json:"sms_type,omitempty" xml:"sms_type,omitempty"`
	// 系统分配的任务code
	TaskCode string `json:"task_code,omitempty" xml:"task_code,omitempty"`
	// 商品或店铺详情页H5长链地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}
