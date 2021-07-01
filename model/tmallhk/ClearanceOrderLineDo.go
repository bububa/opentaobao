package tmallhk

// ClearanceOrderLineDo 结构体
type ClearanceOrderLineDo struct {
	// 货款
	ActualValue int64 `json:"actual_value,omitempty" xml:"actual_value,omitempty"`
	// 商品单价
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 根类目
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 申报要素
	Declaration string `json:"declaration,omitempty" xml:"declaration,omitempty"`
	// 淘系商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品链接
	ItemUrl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 原产国
	OriginalCountry string `json:"original_country,omitempty" xml:"original_country,omitempty"`
	// 付款状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 商品购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 退款状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 主类目
	RootCat string `json:"root_cat,omitempty" xml:"root_cat,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 销售属性
	SellProperty string `json:"sell_property,omitempty" xml:"sell_property,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 税费封装
	TaxDO *ClearanceTaxDo `json:"tax_d_o,omitempty" xml:"tax_d_o,omitempty"`
	// 计量单位封装
	UnitDO *ClearanceUnitDo `json:"unit_d_o,omitempty" xml:"unit_d_o,omitempty"`
	// 镜像商品ID
	VirtualItemId string `json:"virtual_item_id,omitempty" xml:"virtual_item_id,omitempty"`
}
