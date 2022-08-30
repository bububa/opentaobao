package cainiaolocker

// CollectTrackingInfo 结构体
type CollectTrackingInfo struct {
	// 站点id
	StationId string `json:"station_id,omitempty" xml:"station_id,omitempty"`
	// 订单对应的取件人电话
	GetterPhone string `json:"getter_phone,omitempty" xml:"getter_phone,omitempty"`
	// 订单对应的投件人电话
	PostPhone string `json:"post_phone,omitempty" xml:"post_phone,omitempty"`
	// 扩展数据（JSON格式的键值对），如果是取件码取件，请返回取件使用的取件码
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 站点编码
	StationNo string `json:"station_no,omitempty" xml:"station_no,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 动作编码
	ActionCode string `json:"action_code,omitempty" xml:"action_code,omitempty"`
	// 站点订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 快递公司编号
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 订单类型(0-代收业务)
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 动作发生时间戳，单位：毫秒
	ActionTime int64 `json:"action_time,omitempty" xml:"action_time,omitempty"`
}
