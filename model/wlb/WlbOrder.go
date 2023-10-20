package wlb

import (
	"sync"
)

// WlbOrder 结构体
type WlbOrder struct {
	// 出库或者入库，IN表示入库，OUT表示出库
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单来源:&lt;br/&gt;产生物流订单的原因，比如:&lt;br/&gt;&lt;br/&gt;订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 1:正常订单: NORMAL2:退货订单: RETURN3:换货订单: CHANGE
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// (1)其它:    OTHER&lt;br/&gt;(2)淘宝交易: TAOBAO&lt;br/&gt;(3)301:调拨: ALLOCATION&lt;br/&gt;(4)401:盘点:CHECK&lt;br/&gt;(5)501:销售采购:PRUCHASE
	OrderSubType string `json:"order_sub_type,omitempty" xml:"order_sub_type,omitempty"`
	// 卖家NICK
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 发货物流公司编码，STO,YTO,EMS等
	TmsTpCode string `json:"tms_tp_code,omitempty" xml:"tms_tp_code,omitempty"`
	// 物流状态，订单已创建：0订单已取消: -1订单关闭:-3下发中: 10已下发: 11TMS拒签 :-100接收方拒签：-200已发货:100签收成功:200
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 卖家取消,仓库取消标识
	OrderStatusReason string `json:"order_status_reason,omitempty" xml:"order_status_reason,omitempty"`
	// 订单备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 原订单编码
	PrevOrderCode string `json:"prev_order_code,omitempty" xml:"prev_order_code,omitempty"`
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 接收人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 接收人旺旺
	ReceiverWangwang string `json:"receiver_wangwang,omitempty" xml:"receiver_wangwang,omitempty"`
	// 1
	ReceiverMail string `json:"receiver_mail,omitempty" xml:"receiver_mail,omitempty"`
	// 收件人邮编
	ReceiverZipCode string `json:"receiver_zip_code,omitempty" xml:"receiver_zip_code,omitempty"`
	// 接收人手机号码
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 接收人固定电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 发票信息
	InvoiceInfo string `json:"invoice_info,omitempty" xml:"invoice_info,omitempty"`
	// 收件人省份
	ReceiverProvince string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
	// 收件人城市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 区或者县
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
	// 收件人具体地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 发件人所在省份
	SenderProvince string `json:"sender_province,omitempty" xml:"sender_province,omitempty"`
	// 发件人城市
	SenderCity string `json:"sender_city,omitempty" xml:"sender_city,omitempty"`
	// 发件人所在区
	SenderArea string `json:"sender_area,omitempty" xml:"sender_area,omitempty"`
	// 发件人地址
	SenderAddress string `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
	// 发件人姓名
	SenderName string `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
	// 发件人电子邮箱
	SenderEmail string `json:"sender_email,omitempty" xml:"sender_email,omitempty"`
	// 发件人联系电话
	SenderPhone string `json:"sender_phone,omitempty" xml:"sender_phone,omitempty"`
	// 发件人移动电话
	SenderMobile string `json:"sender_mobile,omitempty" xml:"sender_mobile,omitempty"`
	// 发件人邮编
	SenderZipCode string `json:"sender_zip_code,omitempty" xml:"sender_zip_code,omitempty"`
	// 配送开始时间通常是HH:MM格式
	ScheduleDay string `json:"schedule_day,omitempty" xml:"schedule_day,omitempty"`
	// 配送结束时间通常是HH:MM格式
	ScheduleEnd string `json:"schedule_end,omitempty" xml:"schedule_end,omitempty"`
	// 计划送达开始时间
	ExpectStartTime string `json:"expect_start_time,omitempty" xml:"expect_start_time,omitempty"`
	// 计划送达结束时间
	ExpectEndTime string `json:"expect_end_time,omitempty" xml:"expect_end_time,omitempty"`
	// 订单总价
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 第1位:COD,2:限时配送,3:预售,4:需要发票,5:已投诉,第6位:合单,第7位:拆单 第8位：EXCHANGE-换货， 第9位:VISIT-上门 ， 第10位: MODIFYTRANSPORT-是否可改配送方式，第11位：是否物流代理确认发货
	OrderFlag int64 `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// 应收金额
	ReceivableAmount int64 `json:"receivable_amount,omitempty" xml:"receivable_amount,omitempty"`
	// cod服务费
	ServiceFee int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 订单取消状态：&lt;br/&gt;1-取消中； &lt;br/&gt;2-取消失败；&lt;br/&gt;3-取消完成
	CancelOrderStatus int64 `json:"cancel_order_status,omitempty" xml:"cancel_order_status,omitempty"`
	// 发货速度 ， 101-当日达， 102-次晨达， 103-次日达
	ScheduleSpeed int64 `json:"schedule_speed,omitempty" xml:"schedule_speed,omitempty"`
	// 卖家ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 1
	IsCompleted bool `json:"is_completed,omitempty" xml:"is_completed,omitempty"`
}

var poolWlbOrder = sync.Pool{
	New: func() any {
		return new(WlbOrder)
	},
}

// GetWlbOrder() 从对象池中获取WlbOrder
func GetWlbOrder() *WlbOrder {
	return poolWlbOrder.Get().(*WlbOrder)
}

// ReleaseWlbOrder 释放WlbOrder
func ReleaseWlbOrder(v *WlbOrder) {
	v.OperateType = ""
	v.OrderCode = ""
	v.OrderSource = ""
	v.OrderSourceCode = ""
	v.OrderType = ""
	v.OrderSubType = ""
	v.UserNick = ""
	v.StoreCode = ""
	v.TmsTpCode = ""
	v.OrderStatus = ""
	v.OrderStatusReason = ""
	v.Remark = ""
	v.PrevOrderCode = ""
	v.BuyerNick = ""
	v.ReceiverName = ""
	v.ReceiverWangwang = ""
	v.ReceiverMail = ""
	v.ReceiverZipCode = ""
	v.ReceiverMobile = ""
	v.ReceiverPhone = ""
	v.InvoiceInfo = ""
	v.ReceiverProvince = ""
	v.ReceiverCity = ""
	v.ReceiverArea = ""
	v.ReceiverAddress = ""
	v.SenderProvince = ""
	v.SenderCity = ""
	v.SenderArea = ""
	v.SenderAddress = ""
	v.SenderName = ""
	v.SenderEmail = ""
	v.SenderPhone = ""
	v.SenderMobile = ""
	v.SenderZipCode = ""
	v.ScheduleDay = ""
	v.ScheduleEnd = ""
	v.ExpectStartTime = ""
	v.ExpectEndTime = ""
	v.TotalAmount = 0
	v.OrderFlag = 0
	v.ReceivableAmount = 0
	v.ServiceFee = 0
	v.CancelOrderStatus = 0
	v.ScheduleSpeed = 0
	v.UserId = 0
	v.IsCompleted = false
	poolWlbOrder.Put(v)
}
