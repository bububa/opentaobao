package vaccin

// VaccinStoreInfo 结构体
type VaccinStoreInfo struct {
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 地址
	Addr string `json:"addr,omitempty" xml:"addr,omitempty"`
	// 维度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 地区
	Dist string `json:"dist,omitempty" xml:"dist,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// qty
	Qty string `json:"qty,omitempty" xml:"qty,omitempty"`
	// storeName
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// businessTime
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// distN
	DistN string `json:"dist_n,omitempty" xml:"dist_n,omitempty"`
	// storePacket
	StorePacket int64 `json:"store_packet,omitempty" xml:"store_packet,omitempty"`
	// storeId
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
