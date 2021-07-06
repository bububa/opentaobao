package koubeimall

// ItemDto 结构体
type ItemDto struct {
	// 图片相册
	ItemImageList []ItemImage `json:"item_image_list,omitempty" xml:"item_image_list>item_image,omitempty"`
	// 商品类型，包含：TICKET_VOUCHER（购买类卡券类商品），DISH（菜品），BOOK_ITEM（预定类型）
	ItemType string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 商品子类型，包含：TRADE_VOUCHER（套餐购买类商品），VOUCHER_BUY（购买代金券），VOUCHER_PACKAGE（劵包）
	SubItemType string `json:"sub_item_type,omitempty" xml:"sub_item_type,omitempty"`
	// 商品原价
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 商家备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 商品折扣
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 商品售卖价格
	SellPrice string `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// 商品所属门店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商品总销量
	SoldQuantity string `json:"sold_quantity,omitempty" xml:"sold_quantity,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品描述/副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 商品主图/封面图
	ItemCover string `json:"item_cover,omitempty" xml:"item_cover,omitempty"`
	// 商品详情链接，根据入参display_channel信息，获取对应端小程序链接，默认支付宝小程序链接
	ItemDetailUrl string `json:"item_detail_url,omitempty" xml:"item_detail_url,omitempty"`
	// 节省X元，originalPrice - salesPrice
	SavedMoney string `json:"saved_money,omitempty" xml:"saved_money,omitempty"`
	// 描述：节省X元
	SavedMoneyInfo string `json:"saved_money_info,omitempty" xml:"saved_money_info,omitempty"`
	// 销量 “已售xxx份”，销量大于9999："已售XX万份"
	SalesInfo string `json:"sales_info,omitempty" xml:"sales_info,omitempty"`
	// 商品可售库存
	SellableQuantity string `json:"sellable_quantity,omitempty" xml:"sellable_quantity,omitempty"`
	// 聚合最小限购份数，包括每天限购和售卖周期限购,-1表示不限购
	BuyLimit int64 `json:"buy_limit,omitempty" xml:"buy_limit,omitempty"`
	// 商品所属门店信息模型
	ItemStore *ItemStoreDto `json:"item_store,omitempty" xml:"item_store,omitempty"`
}
