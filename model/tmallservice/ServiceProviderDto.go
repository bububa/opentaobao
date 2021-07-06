package tmallservice

// ServiceProviderDto 结构体
type ServiceProviderDto struct {
	// 服务商nick
	TpNick string `json:"tp_nick,omitempty" xml:"tp_nick,omitempty"`
	// 商家网点名称
	SellerStoreName string `json:"seller_store_name,omitempty" xml:"seller_store_name,omitempty"`
	// 商家网点编号
	SellerStoreCode string `json:"seller_store_code,omitempty" xml:"seller_store_code,omitempty"`
	// 服务商id
	TpId int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 商家网点id
	SellerStoreId int64 `json:"seller_store_id,omitempty" xml:"seller_store_id,omitempty"`
	// 服务门店id
	ServiceStoreId int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
}
