package nlife

// RetailItemTopDo 结构体
type RetailItemTopDo struct {
	// sku列表
	SkuList []RetailSkuTopDo `json:"sku_list,omitempty" xml:"sku_list>retail_sku_top_do,omitempty"`
	// supportTaoke
	SupportTaoke bool `json:"support_taoke,omitempty" xml:"support_taoke,omitempty"`
	// supportWzg
	SupportWzg bool `json:"support_wzg,omitempty" xml:"support_wzg,omitempty"`
	// 详情页地址
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// 商品的销售属性
	Props string `json:"props,omitempty" xml:"props,omitempty"`
	// 图片url
	Images string `json:"images,omitempty" xml:"images,omitempty"`
	// 商品主图的url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 价格,单位元,保留2位小数
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品库存
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 货号/款号
	GoodsNo string `json:"goods_no,omitempty" xml:"goods_no,omitempty"`
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品在商家的编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品牌ID
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目ID
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 门店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
	StoreIdType string `json:"store_id_type,omitempty" xml:"store_id_type,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// PC端的商品详情
	PcDesc string `json:"pc_desc,omitempty" xml:"pc_desc,omitempty"`
	// 无线端的商品详情
	WirelessDesc string `json:"wireless_desc,omitempty" xml:"wireless_desc,omitempty"`
	// 主图以外的图片列表
	ImageList []string `json:"image_list,omitempty" xml:"image_list>string,omitempty"`
	// 商品的挂牌价-单位元,保留2位小数
	TagPrice string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
	// 商品的属性名称
	PropsName string `json:"props_name,omitempty" xml:"props_name,omitempty"`
	// 商品类型: 0-IC线上商品; 1-商户导入线下商品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 类目树的信息(自顶向下),格式为 cid1:categoryName1;cid2:categoryName2;cid3:categoryName3
	AllCategoryInfo string `json:"all_category_info,omitempty" xml:"all_category_info,omitempty"`
	// 网直供库存
	OnlineNum int64 `json:"online_num,omitempty" xml:"online_num,omitempty"`
	// 零售商线上商品的itemId
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 供应商的名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// taobaoItemId
	TaobaoItemId int64 `json:"taobao_item_id,omitempty" xml:"taobao_item_id,omitempty"`
}
