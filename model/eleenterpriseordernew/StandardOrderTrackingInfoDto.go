package eleenterpriseordernew

import (
	"sync"
)

// StandardOrderTrackingInfoDto 结构体
type StandardOrderTrackingInfoDto struct {
	// 餐品
	FoodsInfos []FoodsInfo `json:"foods_infos,omitempty" xml:"foods_infos>foods_info,omitempty"`
	// 订单杂项费用
	OrderExtras []OrderExtra `json:"order_extras,omitempty" xml:"order_extras>order_extra,omitempty"`
	// 第三方业务订单编号
	BNo string `json:"b_no,omitempty" xml:"b_no,omitempty"`
	// 序列号（无业务含义）
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	// 饿了么订单Id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 企业支付费用
	EntFee string `json:"ent_fee,omitempty" xml:"ent_fee,omitempty"`
	// 订单备注说明
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 优惠后订单金额
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 成本中心名称
	CostCenterName string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// 税点费用
	TaxFee string `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// 预订时间
	ReceivePlanTime string `json:"receive_plan_time,omitempty" xml:"receive_plan_time,omitempty"`
	// 员工支付金额
	EmployeeFee string `json:"employee_fee,omitempty" xml:"employee_fee,omitempty"`
	// 按钮文字
	ButtonText string `json:"button_text,omitempty" xml:"button_text,omitempty"`
	// 按钮 url
	ButtonUrl string `json:"button_url,omitempty" xml:"button_url,omitempty"`
	// 配送员电话-弃用-兼容旧版本
	DeliveryanPhone string `json:"deliveryan_phone,omitempty" xml:"deliveryan_phone,omitempty"`
	// 配送员电话-弃用-兼容旧版本
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 员工编号
	UNo string `json:"u_no,omitempty" xml:"u_no,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 配送员姓名
	DeliverymanName string `json:"deliveryman_name,omitempty" xml:"deliveryman_name,omitempty"`
	// 配送员电话
	DeliverymanPhone string `json:"deliveryman_phone,omitempty" xml:"deliveryman_phone,omitempty"`
	// 地址信息
	AddressInfo *AddressInfo `json:"address_info,omitempty" xml:"address_info,omitempty"`
	// 餐厅信息
	RestaurantInfo *RestaurantInfo `json:"restaurant_info,omitempty" xml:"restaurant_info,omitempty"`
	// 下单时间（秒）
	CreatedAt int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 订单状态（参考附录）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 配送日期
	DeliveryDate int64 `json:"delivery_date,omitempty" xml:"delivery_date,omitempty"`
	// 1表示展示按钮，0表示不展示按钮
	ShowButton int64 `json:"show_button,omitempty" xml:"show_button,omitempty"`
	// 按钮状态，和button_text一一对应，当show_button为0，button_code也为0，此时button_text为空（参考附录）
	ButtonCode int64 `json:"button_code,omitempty" xml:"button_code,omitempty"`
	// 状态-兼容旧版本
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolStandardOrderTrackingInfoDto = sync.Pool{
	New: func() any {
		return new(StandardOrderTrackingInfoDto)
	},
}

// GetStandardOrderTrackingInfoDto() 从对象池中获取StandardOrderTrackingInfoDto
func GetStandardOrderTrackingInfoDto() *StandardOrderTrackingInfoDto {
	return poolStandardOrderTrackingInfoDto.Get().(*StandardOrderTrackingInfoDto)
}

// ReleaseStandardOrderTrackingInfoDto 释放StandardOrderTrackingInfoDto
func ReleaseStandardOrderTrackingInfoDto(v *StandardOrderTrackingInfoDto) {
	v.FoodsInfos = v.FoodsInfos[:0]
	v.OrderExtras = v.OrderExtras[:0]
	v.BNo = ""
	v.SerialNumber = ""
	v.OrderId = ""
	v.EntFee = ""
	v.Remark = ""
	v.TotalFee = ""
	v.CostCenterName = ""
	v.TaxFee = ""
	v.ReceivePlanTime = ""
	v.EmployeeFee = ""
	v.ButtonText = ""
	v.ButtonUrl = ""
	v.DeliveryanPhone = ""
	v.Extra = ""
	v.UNo = ""
	v.OrderNo = ""
	v.DeliverymanName = ""
	v.DeliverymanPhone = ""
	v.AddressInfo = nil
	v.RestaurantInfo = nil
	v.CreatedAt = 0
	v.Status = 0
	v.DeliveryDate = 0
	v.ShowButton = 0
	v.ButtonCode = 0
	v.StatusCode = 0
	poolStandardOrderTrackingInfoDto.Put(v)
}
