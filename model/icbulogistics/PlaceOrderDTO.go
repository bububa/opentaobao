package icbulogistics

// PlaceOrderDTO 
type PlaceOrderDTO struct {
    // 产品编码
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 1
    NeedPickup   bool `json:"need_pickup,omitempty" xml:"need_pickup,omitempty"`
    // 货品信息
    CargoList   []CargoList `json:"cargo_list,omitempty" xml:"cargo_list>cargo_list,omitempty"`
    // 包裹信息
    PackageList   []PackageList `json:"package_list,omitempty" xml:"package_list>package_list,omitempty"`
    // 收货人地址
    ConsigneeAddress   *ConsigneeAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
    // 发货人地址
    ConsignorAddress   *ConsignorAddress `json:"consignor_address,omitempty" xml:"consignor_address,omitempty"`
    // 申报信息
    ExpressCustoms   *ExpressCustomsDTO `json:"express_customs,omitempty" xml:"express_customs,omitempty"`
    // 信保单ID
    TradeBizId   string `json:"trade_biz_id,omitempty" xml:"trade_biz_id,omitempty"`
    // 发货批次ID
    SupplyChainBizId   string `json:"supply_chain_biz_id,omitempty" xml:"supply_chain_biz_id,omitempty"`
    // 起始地邮编
    OriginZipCode   string `json:"origin_zip_code,omitempty" xml:"origin_zip_code,omitempty"`
    // 目的地邮编
    DestinationZipCode   string `json:"destination_zip_code,omitempty" xml:"destination_zip_code,omitempty"`
    // 交货到仓快递信息
    DeliverWarehouseExpress   *DeliverWarehouseExpressDTO `json:"deliver_warehouse_express,omitempty" xml:"deliver_warehouse_express,omitempty"`
    // 目的地国家
    DestinationCountryCode   string `json:"destination_country_code,omitempty" xml:"destination_country_code,omitempty"`
    // 备用字段（上门揽收地址），目前按发货人地址
    PickupAddress   *ContactAddress `json:"pickup_address,omitempty" xml:"pickup_address,omitempty"`
    // 备用字段（退货地址），目前按退货申请指定地址
    ReturnAddress   *ContactAddress `json:"return_address,omitempty" xml:"return_address,omitempty"`
}
