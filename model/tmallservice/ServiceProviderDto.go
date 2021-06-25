package tmallservice

// ServiceProviderDto 
type ServiceProviderDto struct {

    // 服务商nick
    TpNick   string `json:"tp_nick,omitempty"`

    // 服务商id
    TpId   int64 `json:"tp_id,omitempty"`

    // 商家网点名称
    SellerStoreName   string `json:"seller_store_name,omitempty"`

    // 商家网点编号
    SellerStoreCode   string `json:"seller_store_code,omitempty"`

    // 商家网点id
    SellerStoreId   int64 `json:"seller_store_id,omitempty"`

    // 服务门店id
    ServiceStoreId   int64 `json:"service_store_id,omitempty"`

}
