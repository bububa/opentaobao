package waybill

import (
	"sync"
)

// WaybillApplyFullUpdateRequest 结构体
type WaybillApplyFullUpdateRequest struct {
	// 物流服务能力集合
	LogisticsServiceList []LogisticsService `json:"logistics_service_list,omitempty" xml:"logistics_service_list>logistics_service,omitempty"`
	// 交易订单号（组合表示合并订单）
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
	// 包裹里面的商品名称
	PackageItems []PackageItem `json:"package_items,omitempty" xml:"package_items>package_item,omitempty"`
	// 收件人姓名
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 收件人电话
	ConsigneePhone string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
	// 快递服务商CODE
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递服务产品类型编码
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 发件人姓名
	SendName string `json:"send_name,omitempty" xml:"send_name,omitempty"`
	// 发件人联系方式
	SendPhone string `json:"send_phone,omitempty" xml:"send_phone,omitempty"`
	// 订单渠道类型
	OrderChannelsType string `json:"order_channels_type,omitempty" xml:"order_channels_type,omitempty"`
	// 电子面单单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// ERP 订单号或包裹号
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 使用者ID
	RealUserId int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
	// 包裹重量 单位为G(克)
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 收\发货地址
	ConsigneeAddress *WaybillAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
	// 包裹体积 单位为ML(毫升)或立方厘米
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
}

var poolWaybillApplyFullUpdateRequest = sync.Pool{
	New: func() any {
		return new(WaybillApplyFullUpdateRequest)
	},
}

// GetWaybillApplyFullUpdateRequest() 从对象池中获取WaybillApplyFullUpdateRequest
func GetWaybillApplyFullUpdateRequest() *WaybillApplyFullUpdateRequest {
	return poolWaybillApplyFullUpdateRequest.Get().(*WaybillApplyFullUpdateRequest)
}

// ReleaseWaybillApplyFullUpdateRequest 释放WaybillApplyFullUpdateRequest
func ReleaseWaybillApplyFullUpdateRequest(v *WaybillApplyFullUpdateRequest) {
	v.LogisticsServiceList = v.LogisticsServiceList[:0]
	v.TradeOrderList = v.TradeOrderList[:0]
	v.PackageItems = v.PackageItems[:0]
	v.ConsigneeName = ""
	v.ConsigneePhone = ""
	v.CpCode = ""
	v.ProductType = ""
	v.SendName = ""
	v.SendPhone = ""
	v.OrderChannelsType = ""
	v.WaybillCode = ""
	v.PackageId = ""
	v.RealUserId = 0
	v.Weight = 0
	v.ConsigneeAddress = nil
	v.Volume = 0
	poolWaybillApplyFullUpdateRequest.Put(v)
}
