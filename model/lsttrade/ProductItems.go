package lsttrade

// ProductItems 结构体
type ProductItems struct {
	// cspuID
	CspuId string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
	// 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。
	CargoNumber string `json:"cargo_number,omitempty" xml:"cargo_number,omitempty"`
	// 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考
	ProductSnapshotUrl string `json:"product_snapshot_url,omitempty" xml:"product_snapshot_url,omitempty"`
	// 售卖单位 例如：个、件、箱、吨
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品类型，common：普通品，hg：赠品
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 子订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 发货仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 发货仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 零售通仓库类型。customer：虚仓；cainiao：实仓
	LstWarehouseType string `json:"lst_warehouse_type,omitempty" xml:"lst_warehouse_type,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 本云订单：local，异云订单：remote，如果是实仓订单本字段为空
	VirtualWarehouseType string `json:"virtual_warehouse_type,omitempty" xml:"virtual_warehouse_type,omitempty"`
	// 子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 实付金额，单位分
	ItemAmount int64 `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
	// 原始单价，单位分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// skuID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品的唯一标识
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否组合品
	IsMixSet bool `json:"is_mix_set,omitempty" xml:"is_mix_set,omitempty"`
}
