package drug

// StoreDto 结构体
type StoreDto struct {
	// shopId
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// storeId
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// storeName
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// picURL
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// businessStatus
	BusinessStatus string `json:"business_status,omitempty" xml:"business_status,omitempty"`
	// orderCountDesc
	OrderCountDesc string `json:"order_count_desc,omitempty" xml:"order_count_desc,omitempty"`
	// notice
	Notice string `json:"notice,omitempty" xml:"notice,omitempty"`
	// deliveryTime
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// tele
	Tele string `json:"tele,omitempty" xml:"tele,omitempty"`
	// latitude
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// longitude
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// earliestTimeOfDelivery
	EarliestTimeOfDelivery string `json:"earliest_time_of_delivery,omitempty" xml:"earliest_time_of_delivery,omitempty"`
	// punctualityRate
	PunctualityRate string `json:"punctuality_rate,omitempty" xml:"punctuality_rate,omitempty"`
	// dpavgscore
	Dpavgscore string `json:"dpavgscore,omitempty" xml:"dpavgscore,omitempty"`
	// areas
	Areas string `json:"areas,omitempty" xml:"areas,omitempty"`
	// sellerNick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// deliveryTypeDesc
	DeliveryTypeDesc string `json:"delivery_type_desc,omitempty" xml:"delivery_type_desc,omitempty"`
	// orderCount
	OrderCount int64 `json:"order_count,omitempty" xml:"order_count,omitempty"`
	// minimumAccount
	MinimumAccount int64 `json:"minimum_account,omitempty" xml:"minimum_account,omitempty"`
	// deliveryAccount
	DeliveryAccount int64 `json:"delivery_account,omitempty" xml:"delivery_account,omitempty"`
	// deliveryType
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}
