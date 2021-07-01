package wlb

// WlbProcessStatus 结构体
type WlbProcessStatus struct {
	// 物流宝订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 仓库合作公司编码
	StoreTpCode string `json:"store_tp_code,omitempty" xml:"store_tp_code,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// tms合作公司编码
	TmsTpCode string `json:"tms_tp_code,omitempty" xml:"tms_tp_code,omitempty"`
	// tms合作公司订单编码
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 状态内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 订单操作状态：WMS_ACCEPT;WMS_PRINT;WMS_PICK;WMS_CHECK;WMS_PACKAGE;WMS_CONSIGN;WMS_CANCEL;WMS_UNKNOWN;WMS_CONFIRMEDTMS_ACCEPT;TMS_STATION_IN;TMS_STATION_OUT;TMS_SIGN;TMS_REJECT;TMS_CANCEL;TMS_UNKNOW;SYS_UNKNOWN
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
