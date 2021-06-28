package waybill

// DeliveryStrategyInfo 
/* model for simplify = false
type DeliveryStrategyInfo struct {

    // 识别买家备注: 0-忽略, 1-识别, 2-仅识别合作cp
    
    BuyerMessageRule   int64 `json:"buyer_message_rule,omitempty"`
    

    // 合作CP信息
    
    CocpInfoList  struct {
        CpInfo  []CpInfo `json:"cp_info,omitempty"`
    } `json:"cocp_info_list,omitempty"`
    

    // 特殊线路
    
    SpecialRouteInfoList  struct {
        SpecialRouteInfo  []SpecialRouteInfo `json:"special_route_info,omitempty"`
    } `json:"special_route_info_list,omitempty"`
    

    // 仓id
    
    WarehouseId   int64 `json:"warehouse_id,omitempty"`
    

    // 仓名称
    
    WarehouseName   string `json:"warehouse_name,omitempty"`
    

}
*/

// DeliveryStrategyInfo 
type DeliveryStrategyInfo struct {

    // 识别买家备注: 0-忽略, 1-识别, 2-仅识别合作cp
    BuyerMessageRule   int64 `json:"buyer_message_rule,omitempty"`

    // 合作CP信息
    CocpInfoList   []CpInfo `json:"cocp_info_list,omitempty"`

    // 特殊线路
    SpecialRouteInfoList   []SpecialRouteInfo `json:"special_route_info_list,omitempty"`

    // 仓id
    WarehouseId   int64 `json:"warehouse_id,omitempty"`

    // 仓名称
    WarehouseName   string `json:"warehouse_name,omitempty"`

}
