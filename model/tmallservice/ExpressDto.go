package tmallservice

// ExpressDto 结构体
type ExpressDto struct {
	// 子物流单号
	SubMailNos []string `json:"sub_mail_nos,omitempty" xml:"sub_mail_nos>string,omitempty"`
	// 取件员名称
	CourierName string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
	// 取件员手机号码
	CourierMobile string `json:"courier_mobile,omitempty" xml:"courier_mobile,omitempty"`
	// 寄件单号（废弃）
	ExpressOrderId string `json:"express_order_id,omitempty" xml:"express_order_id,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 物流服务商账号名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 打印免单信息
	PrintInfo string `json:"print_info,omitempty" xml:"print_info,omitempty"`
	// 物品信息明细 字符串格式为：json 串 例子： [{ "name": "大衣,毛衣", "count": 1 }, { "name": "裤子", "count": 2 }]
	GoodsInfo string `json:"goods_info,omitempty" xml:"goods_info,omitempty"`
	// 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 预约开始时间
	ReserveEndTime string `json:"reserve_end_time,omitempty" xml:"reserve_end_time,omitempty"`
	// 预约结束时间
	ReserveStartTime string `json:"reserve_start_time,omitempty" xml:"reserve_start_time,omitempty"`
	// luxury（奢饰商品）/common（普通商品）
	GoodsType string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// 物流订单号
	LogisticsOrderId int64 `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// 物流商账号ID（（ERP服务商依据不同的账号ID，走不同的物流系统对接方式））
	LogisticsTpId int64 `json:"logistics_tp_id,omitempty" xml:"logistics_tp_id,omitempty"`
	// 寄件人信息
	Sender *CustomerInfo `json:"sender,omitempty" xml:"sender,omitempty"`
	// 收件人信息
	Receiver *CustomerInfo `json:"receiver,omitempty" xml:"receiver,omitempty"`
}
