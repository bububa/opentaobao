package tmallnr

// NrLogisticsInfoSynReqDto 结构体
type NrLogisticsInfoSynReqDto struct {
	// 站点名称
	FacilityName string `json:"facility_name,omitempty" xml:"facility_name,omitempty"`
	// 流转节点的详细地址及操作描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 快递员的姓名
	DeliveryUserName string `json:"delivery_user_name,omitempty" xml:"delivery_user_name,omitempty"`
	// 流转节点发生时间
	EventCreateTime string `json:"event_create_time,omitempty" xml:"event_create_time,omitempty"`
	// 取消配送原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 快递公司名称，自配送传入：其他
	MailCpName string `json:"mail_cp_name,omitempty" xml:"mail_cp_name,omitempty"`
	// 流转节点的当前城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 快递单号。各个快递公司的运单号格式不同。
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 取消配送编号
	CancelReasonCode string `json:"cancel_reason_code,omitempty" xml:"cancel_reason_code,omitempty"`
	// 快递业务员联系方式，手机号码或电话。
	DeliveryPhone string `json:"delivery_phone,omitempty" xml:"delivery_phone,omitempty"`
	// 请求编号
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 运单所属CP，自配送传入：OTHER
	MailCp string `json:"mail_cp,omitempty" xml:"mail_cp,omitempty"`
	// 业务类型（oneHour:一小时达,zqs:周期送,dss:定时送，周期送业务:zqs）如果自配送传入：other
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 配送日期，周期送业务必传
	PlanDate string `json:"plan_date,omitempty" xml:"plan_date,omitempty"`
	// 运力类型，0: 第三方快递员,  1:商家自动
	DeliveryUserType int64 `json:"delivery_user_type,omitempty" xml:"delivery_user_type,omitempty"`
	// 快递员id
	DeliveryUserId int64 `json:"delivery_user_id,omitempty" xml:"delivery_user_id,omitempty"`
	// 淘系交易主订单号
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 事件编码,10:已下发等待接单,20:骑手已接单，待提货,40:揽收,999:妥投,50:拒收,-999:取消
	Event int64 `json:"event,omitempty" xml:"event,omitempty"`
	// 操作人类型:1寄件人,3客服小二,4快递员,5CP,6收件人,100系统
	EventOperType int64 `json:"event_oper_type,omitempty" xml:"event_oper_type,omitempty"`
	// 淘宝交易子订单号，周期送业务必传
	SubBizOrderId int64 `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 配送序号
	SequenceNo int64 `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
}
