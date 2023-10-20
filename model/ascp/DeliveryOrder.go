package ascp

import (
	"sync"
)

// DeliveryOrder 结构体
type DeliveryOrder struct {
	// 发票信息
	Invoices []Invoice `json:"invoices,omitempty" xml:"invoices>invoice,omitempty"`
	// 包材信息
	PackageMaterialList []PackageMaterial `json:"package_material_list,omitempty" xml:"package_material_list>package_material,omitempty"`
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 出库单号
	DeliveryOrderCode string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
	// 仓储系统出库单号
	DeliveryOrderId string `json:"delivery_order_id,omitempty" xml:"delivery_order_id,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 出库单类型(JYCK=一般交易出库;HHCK=换货出库;BFCK=补发出库;QTCK=其他出库单)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 出库单状态(NEW-未开始处理;ACCEPT-仓库接单;PARTDELIVERED-部分发货完成;DELIVERED-发货完成;EXCEPTION-异 常;CANCELED-取消;CLOSED-关闭;REJECT-拒单;CANCELEDFAIL-取消失败;只传英文编码)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求 不会被重复处理;条件必填;条件为一单需要多次确认时)
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 支持出库单多次发货(多次发货后确认时;0表示发货单最终状态确认;1表示发货单中间状态确认)
	ConfirmType string `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 订单完成时间(YYYY-MM-DD HH:MM:SS)
	OrderConfirmTime string `json:"order_confirm_time,omitempty" xml:"order_confirm_time,omitempty"`
	// 当前状态操作员编码
	OperatorCode string `json:"operator_code,omitempty" xml:"operator_code,omitempty"`
	// 当前状态操作员姓名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 当前状态操作时间(YYYY-MM-DD HH:MM:SS)
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 仓储费用
	StorageFee string `json:"storage_fee,omitempty" xml:"storage_fee,omitempty"`
	// 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通 (ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 物流公司名称
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 包裹编号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 包裹长度(单位：厘米)
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度(单位：厘米)
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高度(单位：厘米)
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹理论重量(单位：千克)
	TheoreticalWeight string `json:"theoretical_weight,omitempty" xml:"theoretical_weight,omitempty"`
	// 包裹重量(单位：千克)
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹体积(单位：升)
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 发票号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 原出库单号(ERP分配)
	PreDeliveryOrderCode string `json:"pre_delivery_order_code,omitempty" xml:"pre_delivery_order_code,omitempty"`
	// 原出库单号(WMS分配)
	PreDeliveryOrderId string `json:"pre_delivery_order_id,omitempty" xml:"pre_delivery_order_id,omitempty"`
	// 订单标记(用字符串格式来表示订单标记列表:例如COD=货到付款;LIMIT=限时配 送;PRESELL=预 售;COMPLAIN=已投诉;SPLIT=拆单;EXCHANGE=换货;VISIT=上门;MODIFYTRANSPORT=是否 可改配送方式;CONSIGN = 物流宝代理发货;SELLER_AFFORD=是否卖家承担运费;FENXIAO=分销订 单)
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// 订单来源平台编码(TB=淘宝、TM=天猫、MAOCHAO =天猫超市、TMGJZY = 天猫国际直营、PTTMKLDJK = 天猫国际、TMYPDQ = 天猫优品、JD=京东、DD=当当、PP=拍拍、YX= 易讯、 EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂 、MGJ=蘑菇街、 JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、 MIA=蜜芽、KLZY = 猫享&amp;考拉自营、OTHER=其他(只传英文编码))
	SourcePlatformCode string `json:"source_platform_code,omitempty" xml:"source_platform_code,omitempty"`
	// 订单来源平台名称
	SourcePlatformName string `json:"source_platform_name,omitempty" xml:"source_platform_name,omitempty"`
	// 发货单创建时间(YYYY-MM-DD HH:MM:SS)
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 前台订单/店铺订单的创建时间/下单时间
	PlaceOrderTime string `json:"place_order_time,omitempty" xml:"place_order_time,omitempty"`
	// 订单支付时间(YYYY-MM-DD HH:MM:SS)
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付平台交易号
	PayNo string `json:"pay_no,omitempty" xml:"pay_no,omitempty"`
	// 店铺名称
	ShopNick string `json:"shop_nick,omitempty" xml:"shop_nick,omitempty"`
	// 卖家名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单总金额(订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用 ;单位 元)
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 商品总金额(元)
	ItemAmount string `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
	// 订单折扣金额(元)
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 快递费用(元)
	Freight string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 应收金额(消费者还需要支付多少--货到付款时消费者还需要支付多少约定使用这个字 段;单位元 )
	ArAmount string `json:"ar_amount,omitempty" xml:"ar_amount,omitempty"`
	// 已收金额(消费者已经支付多少;单位元)
	GotAmount string `json:"got_amount,omitempty" xml:"got_amount,omitempty"`
	// COD服务费
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 快递区域编码
	LogisticsAreaCode string `json:"logistics_area_code,omitempty" xml:"logistics_area_code,omitempty"`
	// 是否紧急(Y/N;默认为N)
	IsUrgency string `json:"is_urgency,omitempty" xml:"is_urgency,omitempty"`
	// 是否需要发票(Y/N;默认为N)
	InvoiceFlag string `json:"invoice_flag,omitempty" xml:"invoice_flag,omitempty"`
	// 是否需要保险(Y/N;默认为N)
	InsuranceFlag string `json:"insurance_flag,omitempty" xml:"insurance_flag,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 卖家留言
	SellerMessage string `json:"seller_message,omitempty" xml:"seller_message,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 旧版本货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 最晚揽收时间, string (19) , YYYY-MM-DD HH:MM:SS
	LatestCollectionTime string `json:"latest_collection_time,omitempty" xml:"latest_collection_time,omitempty"`
	// 最晚发货时间, string (19) , YYYY-MM-DD HH:MM:SS
	LatestDeliveryTime string `json:"latest_delivery_time,omitempty" xml:"latest_delivery_time,omitempty"`
	// 交易平台订单
	OaidOrderSourceCode string `json:"oaid_order_source_code,omitempty" xml:"oaid_order_source_code,omitempty"`
	// 单据号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓储系统单据号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 发货要求列表
	DeliveryRequirements *DeliveryRequirement `json:"delivery_requirements,omitempty" xml:"delivery_requirements,omitempty"`
	// 发件人信息
	SenderInfo *SenderInfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 收件人信息
	ReceiverInfo *ReceiverInfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 保险信息
	Insurance *Insurance `json:"insurance,omitempty" xml:"insurance,omitempty"`
}

var poolDeliveryOrder = sync.Pool{
	New: func() any {
		return new(DeliveryOrder)
	},
}

// GetDeliveryOrder() 从对象池中获取DeliveryOrder
func GetDeliveryOrder() *DeliveryOrder {
	return poolDeliveryOrder.Get().(*DeliveryOrder)
}

// ReleaseDeliveryOrder 释放DeliveryOrder
func ReleaseDeliveryOrder(v *DeliveryOrder) {
	v.Invoices = v.Invoices[:0]
	v.PackageMaterialList = v.PackageMaterialList[:0]
	v.Items = v.Items[:0]
	v.DeliveryOrderCode = ""
	v.DeliveryOrderId = ""
	v.WarehouseCode = ""
	v.OrderType = ""
	v.Status = ""
	v.OutBizCode = ""
	v.ConfirmType = ""
	v.OrderConfirmTime = ""
	v.OperatorCode = ""
	v.OperatorName = ""
	v.OperateTime = ""
	v.StorageFee = ""
	v.LogisticsCode = ""
	v.LogisticsName = ""
	v.ExpressCode = ""
	v.PackageCode = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	v.TheoreticalWeight = ""
	v.Weight = ""
	v.Volume = ""
	v.InvoiceNo = ""
	v.PreDeliveryOrderCode = ""
	v.PreDeliveryOrderId = ""
	v.OrderFlag = ""
	v.SourcePlatformCode = ""
	v.SourcePlatformName = ""
	v.CreateTime = ""
	v.PlaceOrderTime = ""
	v.PayTime = ""
	v.PayNo = ""
	v.ShopNick = ""
	v.SellerNick = ""
	v.BuyerNick = ""
	v.TotalAmount = ""
	v.ItemAmount = ""
	v.DiscountAmount = ""
	v.Freight = ""
	v.ArAmount = ""
	v.GotAmount = ""
	v.ServiceFee = ""
	v.LogisticsAreaCode = ""
	v.IsUrgency = ""
	v.InvoiceFlag = ""
	v.InsuranceFlag = ""
	v.BuyerMessage = ""
	v.SellerMessage = ""
	v.Remark = ""
	v.ServiceCode = ""
	v.OwnerCode = ""
	v.LatestCollectionTime = ""
	v.LatestDeliveryTime = ""
	v.OaidOrderSourceCode = ""
	v.OrderCode = ""
	v.OrderId = ""
	v.DeliveryRequirements = nil
	v.SenderInfo = nil
	v.ReceiverInfo = nil
	v.Insurance = nil
	poolDeliveryOrder.Put(v)
}
