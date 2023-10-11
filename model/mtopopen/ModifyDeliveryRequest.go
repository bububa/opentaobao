package mtopopen

// ModifyDeliveryRequest 结构体
type ModifyDeliveryRequest struct {
	// 订单ID
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 快递公司标准编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 收货方式对应的参数，送货上门：{&#34;day&#34;:&#34;2023-03-22&#34;,&#34;times&#34;:&#34;10:00-13:00&#34;} 派送到指定地点：{&#34;point&#34;:&#34;物业&#34;} 找人代收：{&#34;fullName&#34;:&#34;xxxx&#34;,&#34;mobile&#34;:&#34;xxxxxx&#34;}
	ExtendParam string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
	// 修改渠道-淘宝小程序：taoBao，非淘宝：others
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 收货方式「送货上门：0，快递柜/代收点：1，指定地点：2，找人代收：3，其他：-1」
	ReceiveType int64 `json:"receive_type,omitempty" xml:"receive_type,omitempty"`
}
