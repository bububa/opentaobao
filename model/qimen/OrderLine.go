package qimen

import (
	"sync"
)

// OrderLine 结构体
type OrderLine struct {
	// 批次列表
	Batchs []TaobaoQimenDeliveryorderBatchconfirmBatch `json:"batchs,omitempty" xml:"batchs>taobao_qimen_deliveryorder_batchconfirm_batch,omitempty"`
	// 单据行号
	OrderLineNo string `json:"orderLineNo,omitempty" xml:"orderLineNo,omitempty"`
	// 平台交易订单编码
	OrderSourceCode string `json:"orderSourceCode,omitempty" xml:"orderSourceCode,omitempty"`
	// 平台交易子订单编码
	SubSourceCode string `json:"subSourceCode,omitempty" xml:"subSourceCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 商品编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 商品仓储系统编码
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 商品名称
	ItemName string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 交易平台商品编码
	ExtCode string `json:"extCode,omitempty" xml:"extCode,omitempty"`
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 商品的二维码(类似电子产品的SN码;用来进行商品的溯源;多个二维码之间用分号;隔开)
	QrCode string `json:"qrCode,omitempty" xml:"qrCode,omitempty"`
	// 仓库拆单子发货单号
	SubDeliveryOrderId string `json:"subDeliveryOrderId,omitempty" xml:"subDeliveryOrderId,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
	// 交易平台订单
	SourceOrderCode string `json:"sourceOrderCode,omitempty" xml:"sourceOrderCode,omitempty"`
	// 交易平台子订单编码
	SubSourceOrderCode string `json:"subSourceOrderCode,omitempty" xml:"subSourceOrderCode,omitempty"`
	// 支付平台交易号(淘系订单传支付宝交易号)
	PayNo string `json:"payNo,omitempty" xml:"payNo,omitempty"`
	// 零售价(零售价=实际成交价+单件商品折扣金额)
	RetailPrice string `json:"retailPrice,omitempty" xml:"retailPrice,omitempty"`
	// 实际成交价
	ActualPrice string `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// 单件商品折扣金额
	DiscountAmount string `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货品sn编码
	SnCode string `json:"snCode,omitempty" xml:"snCode,omitempty"`
	// 外部业务编码(消息ID;用于去重;当单据需要分批次发送时使用)
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品属性
	SkuProperty string `json:"skuProperty,omitempty" xml:"skuProperty,omitempty"`
	// 采购价
	PurchasePrice string `json:"purchasePrice,omitempty" xml:"purchasePrice,omitempty"`
	// 出库单号, string (50) , 必填
	DeliveryOrderCode string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
	// 奇门仓储字段
	ProductCode string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// 奇门仓储字段
	StockInQty string `json:"stockInQty,omitempty" xml:"stockInQty,omitempty"`
	// 奇门仓储字段
	StockOutQty string `json:"stockOutQty,omitempty" xml:"stockOutQty,omitempty"`
	// 奇门仓储字段
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 奇门仓储字段
	DeliveryOrderId string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
	// 奇门仓储字段
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 奇门仓储字段
	TaobaoItemCode string `json:"taobaoItemCode,omitempty" xml:"taobaoItemCode,omitempty"`
	// 奇门仓储字段
	DiscountPrice string `json:"discountPrice,omitempty" xml:"discountPrice,omitempty"`
	// 奇门仓储字段
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 奇门仓储字段
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// 奇门仓储字段
	StandardPrice string `json:"standardPrice,omitempty" xml:"standardPrice,omitempty"`
	// 奇门仓储字段
	ReferencePrice string `json:"referencePrice,omitempty" xml:"referencePrice,omitempty"`
	// 奇门仓储字段
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 奇门仓储字段
	StandardAmount string `json:"standardAmount,omitempty" xml:"standardAmount,omitempty"`
	// 奇门仓储字段
	SettlementAmount string `json:"settlementAmount,omitempty" xml:"settlementAmount,omitempty"`
	// 奇门仓储字段
	LocationCode string `json:"locationCode,omitempty" xml:"locationCode,omitempty"`
	// 奇门仓储字段
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 奇门仓储字段
	MoveOutLocation string `json:"moveOutLocation,omitempty" xml:"moveOutLocation,omitempty"`
	// 奇门仓储字段
	MoveInLocation string `json:"moveInLocation,omitempty" xml:"moveInLocation,omitempty"`
	// 奇门仓储字段
	ExceptionQty string `json:"exceptionQty,omitempty" xml:"exceptionQty,omitempty"`
	// 用字符串格式来表示订单标记列表(比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL等;中间用“^”来隔开订单标记list (所 有字母全部大写) VISIT=上门；SELLER_AFFORD=是否卖家承担运费(默认是)SYNC_RETURN_BILL=同时退回发票)
	OrderFlag string `json:"orderFlag,omitempty" xml:"orderFlag,omitempty"`
	// 退货原因
	ReturnReason string `json:"returnReason,omitempty" xml:"returnReason,omitempty"`
	// 交易平台商品编码
	PlatformCode string `json:"platformCode,omitempty" xml:"platformCode,omitempty"`
	// 应发商品数量
	PlanQty int64 `json:"planQty,omitempty" xml:"planQty,omitempty"`
	// 实发商品数量
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// snList
	SnList *SnList `json:"snList,omitempty" xml:"snList,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockoutCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolOrderLine = sync.Pool{
	New: func() any {
		return new(OrderLine)
	},
}

// GetOrderLine() 从对象池中获取OrderLine
func GetOrderLine() *OrderLine {
	return poolOrderLine.Get().(*OrderLine)
}

// ReleaseOrderLine 释放OrderLine
func ReleaseOrderLine(v *OrderLine) {
	v.Batchs = v.Batchs[:0]
	v.OrderLineNo = ""
	v.OrderSourceCode = ""
	v.SubSourceCode = ""
	v.OwnerCode = ""
	v.ItemCode = ""
	v.ItemId = ""
	v.InventoryType = ""
	v.ItemName = ""
	v.ExtCode = ""
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.QrCode = ""
	v.SubDeliveryOrderId = ""
	v.SupplierCode = ""
	v.SupplierName = ""
	v.SourceOrderCode = ""
	v.SubSourceOrderCode = ""
	v.PayNo = ""
	v.RetailPrice = ""
	v.ActualPrice = ""
	v.DiscountAmount = ""
	v.Quantity = ""
	v.SnCode = ""
	v.OutBizCode = ""
	v.Remark = ""
	v.Unit = ""
	v.SkuProperty = ""
	v.PurchasePrice = ""
	v.DeliveryOrderCode = ""
	v.ProductCode = ""
	v.StockInQty = ""
	v.StockOutQty = ""
	v.WarehouseCode = ""
	v.DeliveryOrderId = ""
	v.Status = ""
	v.TaobaoItemCode = ""
	v.DiscountPrice = ""
	v.Color = ""
	v.Size = ""
	v.StandardPrice = ""
	v.ReferencePrice = ""
	v.Discount = ""
	v.StandardAmount = ""
	v.SettlementAmount = ""
	v.LocationCode = ""
	v.Amount = ""
	v.MoveOutLocation = ""
	v.MoveInLocation = ""
	v.ExceptionQty = ""
	v.OrderFlag = ""
	v.ReturnReason = ""
	v.PlatformCode = ""
	v.PlanQty = 0
	v.ActualQty = 0
	v.SnList = nil
	v.ExtendProps = nil
	poolOrderLine.Put(v)
}
