package traveltrade

// ProductUpdatePushDto 结构体
type ProductUpdatePushDto struct {
	// 日历价格库存信息  日历价格库存信息
	PriceStocks []ProductPriceStockDto `json:"price_stocks,omitempty" xml:"price_stocks>product_price_stock_dto,omitempty"`
	// 系统商商品码
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 床型ID
	BedId string `json:"bed_id,omitempty" xml:"bed_id,omitempty"`
	// 酒店ID
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 房型ID
	RoomId string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 产品变更通知类型 1：价格，2：库存，3：价库
	NotifyType int64 `json:"notify_type,omitempty" xml:"notify_type,omitempty"`
	// 模式 默认值1；1:普通日历/预约商品（非通兑和非任选）
	Mode int64 `json:"mode,omitempty" xml:"mode,omitempty"`
}
