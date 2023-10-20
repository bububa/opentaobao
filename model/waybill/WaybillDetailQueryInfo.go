package waybill

import (
	"sync"
)

// WaybillDetailQueryInfo 结构体
type WaybillDetailQueryInfo struct {
	// 交易订单列表
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
	// 物流服务能力集合
	LogisticsServiceList []LogisticsService `json:"logistics_service_list,omitempty" xml:"logistics_service_list>logistics_service,omitempty"`
	// 包裹里面的商品类型
	PackageItems []PackageItem `json:"package_items,omitempty" xml:"package_items>package_item,omitempty"`
	// 包裹对应的派件（收件）物流服务商网点（分支机构）代码
	ConsigneeBranchCode string `json:"consignee_branch_code,omitempty" xml:"consignee_branch_code,omitempty"`
	// 包裹对应的派件（收件）物流服务商网点（分支机构）名称
	ConsigneeBranchName string `json:"consignee_branch_name,omitempty" xml:"consignee_branch_name,omitempty"`
	// 收件人姓名
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 收件人联系方式
	ConsigneePhone string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
	// 物流商编码CODE
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后一次打印时间
	LastPrintTime string `json:"last_print_time,omitempty" xml:"last_print_time,omitempty"`
	// 集包地、目的地中心代码。打印时根据该 code 生成目的地中心的条码，条码生成的算法与对应的电子面单条码一致
	PackageCenterCode string `json:"package_center_code,omitempty" xml:"package_center_code,omitempty"`
	// 集包地、目的地中心名称
	PackageCenterName string `json:"package_center_name,omitempty" xml:"package_center_name,omitempty"`
	// ERP订单号或包裹号
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 揽收时间
	PickupTime string `json:"pickup_time,omitempty" xml:"pickup_time,omitempty"`
	// 打印配置项
	PrintConfig string `json:"print_config,omitempty" xml:"print_config,omitempty"`
	// 快递服务产品类型编码
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 发件人姓名
	SendName string `json:"send_name,omitempty" xml:"send_name,omitempty"`
	// 发件人联系方式
	SendPhone string `json:"send_phone,omitempty" xml:"send_phone,omitempty"`
	// 发货网点编码
	ShippingBranchCode string `json:"shipping_branch_code,omitempty" xml:"shipping_branch_code,omitempty"`
	// 发货网点信息
	ShippingBranchName string `json:"shipping_branch_name,omitempty" xml:"shipping_branch_name,omitempty"`
	// 大头笔信息
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 电子面单信息
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 收货人地址
	ConsigneeAddress *WaybillAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
	// 打印次数
	PrintCount int64 `json:"print_count,omitempty" xml:"print_count,omitempty"`
	// 发货地址
	ShippingAddress *WaybillAddress `json:"shipping_address,omitempty" xml:"shipping_address,omitempty"`
	// 面单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 使用者ID
	RealUserId int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
	// 包裹重量 单位为G(克)
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 包裹体积 单位为ML(毫升)或立方厘米
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolWaybillDetailQueryInfo = sync.Pool{
	New: func() any {
		return new(WaybillDetailQueryInfo)
	},
}

// GetWaybillDetailQueryInfo() 从对象池中获取WaybillDetailQueryInfo
func GetWaybillDetailQueryInfo() *WaybillDetailQueryInfo {
	return poolWaybillDetailQueryInfo.Get().(*WaybillDetailQueryInfo)
}

// ReleaseWaybillDetailQueryInfo 释放WaybillDetailQueryInfo
func ReleaseWaybillDetailQueryInfo(v *WaybillDetailQueryInfo) {
	v.TradeOrderList = v.TradeOrderList[:0]
	v.LogisticsServiceList = v.LogisticsServiceList[:0]
	v.PackageItems = v.PackageItems[:0]
	v.ConsigneeBranchCode = ""
	v.ConsigneeBranchName = ""
	v.ConsigneeName = ""
	v.ConsigneePhone = ""
	v.CpCode = ""
	v.CreateTime = ""
	v.LastPrintTime = ""
	v.PackageCenterCode = ""
	v.PackageCenterName = ""
	v.PackageId = ""
	v.PickupTime = ""
	v.PrintConfig = ""
	v.ProductType = ""
	v.SendName = ""
	v.SendPhone = ""
	v.ShippingBranchCode = ""
	v.ShippingBranchName = ""
	v.ShortAddress = ""
	v.SignTime = ""
	v.WaybillCode = ""
	v.ConsigneeAddress = nil
	v.PrintCount = 0
	v.ShippingAddress = nil
	v.Status = 0
	v.RealUserId = 0
	v.Volume = 0
	v.Weight = 0
	poolWaybillDetailQueryInfo.Put(v)
}
