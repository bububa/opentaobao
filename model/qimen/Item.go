package qimen

import (
	"sync"
)

// Item 结构体
type Item struct {
	// test
	Batchs []TaobaoQimenCombineitemSynchronizeBatch `json:"batchs,omitempty" xml:"batchs>taobao_qimen_combineitem_synchronize_batch,omitempty"`
	// 商品编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 后端商品编码
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// ownerCode
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// test
	BrandName string `json:"brandName,omitempty" xml:"brandName,omitempty"`
	// test
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// test
	IsSNMgmt string `json:"isSNMgmt,omitempty" xml:"isSNMgmt,omitempty"`
	// test
	BarCode string `json:"barCode,omitempty" xml:"barCode,omitempty"`
	// test
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// test
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// test
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// test
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// test
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// test
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// test
	GrossWeight string `json:"grossWeight,omitempty" xml:"grossWeight,omitempty"`
	// test
	NetWeight string `json:"netWeight,omitempty" xml:"netWeight,omitempty"`
	// test
	TareWeight string `json:"tareWeight,omitempty" xml:"tareWeight,omitempty"`
	// test
	SafetyStock string `json:"safetyStock,omitempty" xml:"safetyStock,omitempty"`
	// test
	StockUnit string `json:"stockUnit,omitempty" xml:"stockUnit,omitempty"`
	// test
	StockStatus string `json:"stockStatus,omitempty" xml:"stockStatus,omitempty"`
	// test
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// test
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// test
	IsShelfLifeMgmt string `json:"isShelfLifeMgmt,omitempty" xml:"isShelfLifeMgmt,omitempty"`
	// test
	ShelfLife string `json:"shelfLife,omitempty" xml:"shelfLife,omitempty"`
	// test
	RejectLifecycle string `json:"rejectLifecycle,omitempty" xml:"rejectLifecycle,omitempty"`
	// test
	LockupLifecycle string `json:"lockupLifecycle,omitempty" xml:"lockupLifecycle,omitempty"`
	// test
	AdventLifecycle string `json:"adventLifecycle,omitempty" xml:"adventLifecycle,omitempty"`
	// test
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// test
	BatchRemark string `json:"batchRemark,omitempty" xml:"batchRemark,omitempty"`
	// test
	IsBatchMgmt string `json:"isBatchMgmt,omitempty" xml:"isBatchMgmt,omitempty"`
	// test
	PackCode string `json:"packCode,omitempty" xml:"packCode,omitempty"`
	// test
	Pcs string `json:"pcs,omitempty" xml:"pcs,omitempty"`
	// test
	OriginAddress string `json:"originAddress,omitempty" xml:"originAddress,omitempty"`
	// test
	ApprovalNumber string `json:"approvalNumber,omitempty" xml:"approvalNumber,omitempty"`
	// test
	IsFragile string `json:"isFragile,omitempty" xml:"isFragile,omitempty"`
	// test
	IsHazardous string `json:"isHazardous,omitempty" xml:"isHazardous,omitempty"`
	// test
	PricingCategory string `json:"pricingCategory,omitempty" xml:"pricingCategory,omitempty"`
	// test
	IsSku string `json:"isSku,omitempty" xml:"isSku,omitempty"`
	// test
	PackageMaterial string `json:"packageMaterial,omitempty" xml:"packageMaterial,omitempty"`
	// test
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// test
	IsAreaSale string `json:"isAreaSale,omitempty" xml:"isAreaSale,omitempty"`
	// test
	NormalQty string `json:"normalQty,omitempty" xml:"normalQty,omitempty"`
	// test
	DefectiveQty string `json:"defectiveQty,omitempty" xml:"defectiveQty,omitempty"`
	// test
	ReceiveQty string `json:"receiveQty,omitempty" xml:"receiveQty,omitempty"`
	// test
	ExCode string `json:"exCode,omitempty" xml:"exCode,omitempty"`
	// test
	DiscountPrice string `json:"discountPrice,omitempty" xml:"discountPrice,omitempty"`
	// test
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// test
	PlanQty string `json:"planQty,omitempty" xml:"planQty,omitempty"`
	// test
	SourceOrderCode string `json:"sourceOrderCode,omitempty" xml:"sourceOrderCode,omitempty"`
	// test
	SubSourceOrderCode string `json:"subSourceOrderCode,omitempty" xml:"subSourceOrderCode,omitempty"`
	// test
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// test
	OrderLineNo string `json:"orderLineNo,omitempty" xml:"orderLineNo,omitempty"`
	// test
	ActualQty string `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// test
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// test
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// test
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// test
	LockQuantity string `json:"lockQuantity,omitempty" xml:"lockQuantity,omitempty"`
	// test
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// test
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// test
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// test
	ProductCode string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// test
	PaperQty string `json:"paperQty,omitempty" xml:"paperQty,omitempty"`
	// test
	DiffQuantity string `json:"diffQuantity,omitempty" xml:"diffQuantity,omitempty"`
	// test
	ExtCode string `json:"extCode,omitempty" xml:"extCode,omitempty"`
	// test
	LackQty string `json:"lackQty,omitempty" xml:"lackQty,omitempty"`
	// test
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// test
	SnCode string `json:"snCode,omitempty" xml:"snCode,omitempty"`
	// test
	GoodsCode string `json:"goodsCode,omitempty" xml:"goodsCode,omitempty"`
	// test
	StandardPrice string `json:"standardPrice,omitempty" xml:"standardPrice,omitempty"`
	// test
	ReferencePrice string `json:"referencePrice,omitempty" xml:"referencePrice,omitempty"`
	// test
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// test
	ActualAmount string `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	// test
	LatestUpdateTime string `json:"latestUpdateTime,omitempty" xml:"latestUpdateTime,omitempty"`
	// test
	ChangeTime string `json:"changeTime,omitempty" xml:"changeTime,omitempty"`
	// test
	TempRequirement string `json:"tempRequirement,omitempty" xml:"tempRequirement,omitempty"`
	// test
	ChannelCode string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// test
	OriginCode string `json:"originCode,omitempty" xml:"originCode,omitempty"`
	// 商品名称
	ItemName string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 商品简称
	ShortName string `json:"shortName,omitempty" xml:"shortName,omitempty"`
	// 英文名
	EnglishName string `json:"englishName,omitempty" xml:"englishName,omitempty"`
	// 商品属性(如红色;XXL)
	SkuProperty string `json:"skuProperty,omitempty" xml:"skuProperty,omitempty"`
	// 渠道中的商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品类别ID
	CategoryId string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// 商品类别名称
	CategoryName string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// 商品类型(ZC=正常商品;FX=分销商品;ZH=组合商品;ZP=赠品;BC=包材;HC=耗材;FL=辅料;XN=虚拟品;FS=附属品;CC=残次品;OTHER=其它;只传英文编码)
	ItemType string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 吊牌价
	TagPrice string `json:"tagPrice,omitempty" xml:"tagPrice,omitempty"`
	// 零售价
	RetailPrice string `json:"retailPrice,omitempty" xml:"retailPrice,omitempty"`
	// 成本价
	CostPrice string `json:"costPrice,omitempty" xml:"costPrice,omitempty"`
	// 采购价
	PurchasePrice string `json:"purchasePrice,omitempty" xml:"purchasePrice,omitempty"`
	// 季节编码
	SeasonCode string `json:"seasonCode,omitempty" xml:"seasonCode,omitempty"`
	// 季节名称
	SeasonName string `json:"seasonName,omitempty" xml:"seasonName,omitempty"`
	// 品牌代码
	BrandCode string `json:"brandCode,omitempty" xml:"brandCode,omitempty"`
	// 创建时间(YYYY-MM-DD HH:MM:SS)
	CreateTime string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 更新时间(YYYY-MM-DD HH:MM:SS)
	UpdateTime string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 是否有效(Y/N ;默认为N)
	IsValid string `json:"isValid,omitempty" xml:"isValid,omitempty"`
	// 销售配送方式（0=自配|1=菜鸟）
	LogisticsType string `json:"logisticsType,omitempty" xml:"logisticsType,omitempty"`
	// 是否液体, Y/N,  (默认为N)
	IsLiquid string `json:"isLiquid,omitempty" xml:"isLiquid,omitempty"`
	// 签收状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// temp
	SupplierCode string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// isLocked
	IsLocked string `json:"isLocked,omitempty" xml:"isLocked,omitempty"`
	// 货品编码,HZ1234,string(50),,
	ScItemCode string `json:"scItemCode,omitempty" xml:"scItemCode,omitempty"`
	// 计划调拨数量,Item1234,string(50),,
	PlanCount string `json:"planCount,omitempty" xml:"planCount,omitempty"`
	// 实际出库数量,Item1234,string(50),,
	OutCount string `json:"outCount,omitempty" xml:"outCount,omitempty"`
	// 实际入库数量,Item1234,string(50),,
	InCount string `json:"inCount,omitempty" xml:"inCount,omitempty"`
	// 组合商品中的该商品个数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// test
	PriceAdjustment *PriceAdjustment `json:"priceAdjustment,omitempty" xml:"priceAdjustment,omitempty"`
	// 库存商品总量
	TotalQty int64 `json:"totalQty,omitempty" xml:"totalQty,omitempty"`
}

var poolItem = sync.Pool{
	New: func() any {
		return new(Item)
	},
}

// GetItem() 从对象池中获取Item
func GetItem() *Item {
	return poolItem.Get().(*Item)
}

// ReleaseItem 释放Item
func ReleaseItem(v *Item) {
	v.Batchs = v.Batchs[:0]
	v.ItemCode = ""
	v.ItemId = ""
	v.OwnerCode = ""
	v.BrandName = ""
	v.Sn = ""
	v.IsSNMgmt = ""
	v.BarCode = ""
	v.Color = ""
	v.Size = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	v.Volume = ""
	v.GrossWeight = ""
	v.NetWeight = ""
	v.TareWeight = ""
	v.SafetyStock = ""
	v.StockUnit = ""
	v.StockStatus = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.IsShelfLifeMgmt = ""
	v.ShelfLife = ""
	v.RejectLifecycle = ""
	v.LockupLifecycle = ""
	v.AdventLifecycle = ""
	v.BatchCode = ""
	v.BatchRemark = ""
	v.IsBatchMgmt = ""
	v.PackCode = ""
	v.Pcs = ""
	v.OriginAddress = ""
	v.ApprovalNumber = ""
	v.IsFragile = ""
	v.IsHazardous = ""
	v.PricingCategory = ""
	v.IsSku = ""
	v.PackageMaterial = ""
	v.Price = ""
	v.IsAreaSale = ""
	v.NormalQty = ""
	v.DefectiveQty = ""
	v.ReceiveQty = ""
	v.ExCode = ""
	v.DiscountPrice = ""
	v.InventoryType = ""
	v.PlanQty = ""
	v.SourceOrderCode = ""
	v.SubSourceOrderCode = ""
	v.ProduceCode = ""
	v.OrderLineNo = ""
	v.ActualQty = ""
	v.Amount = ""
	v.Unit = ""
	v.WarehouseCode = ""
	v.LockQuantity = ""
	v.OrderCode = ""
	v.OrderType = ""
	v.OutBizCode = ""
	v.ProductCode = ""
	v.PaperQty = ""
	v.DiffQuantity = ""
	v.ExtCode = ""
	v.LackQty = ""
	v.Reason = ""
	v.SnCode = ""
	v.GoodsCode = ""
	v.StandardPrice = ""
	v.ReferencePrice = ""
	v.Discount = ""
	v.ActualAmount = ""
	v.LatestUpdateTime = ""
	v.ChangeTime = ""
	v.TempRequirement = ""
	v.ChannelCode = ""
	v.OriginCode = ""
	v.ItemName = ""
	v.Remark = ""
	v.ShortName = ""
	v.EnglishName = ""
	v.SkuProperty = ""
	v.Title = ""
	v.CategoryId = ""
	v.CategoryName = ""
	v.ItemType = ""
	v.TagPrice = ""
	v.RetailPrice = ""
	v.CostPrice = ""
	v.PurchasePrice = ""
	v.SeasonCode = ""
	v.SeasonName = ""
	v.BrandCode = ""
	v.CreateTime = ""
	v.UpdateTime = ""
	v.IsValid = ""
	v.LogisticsType = ""
	v.IsLiquid = ""
	v.Status = ""
	v.SupplierCode = ""
	v.IsLocked = ""
	v.ScItemCode = ""
	v.PlanCount = ""
	v.OutCount = ""
	v.InCount = ""
	v.Quantity = 0
	v.PriceAdjustment = nil
	v.TotalQty = 0
	poolItem.Put(v)
}
