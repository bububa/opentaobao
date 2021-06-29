package ascpchannel

// Transferdetaildtolist 
type Transferdetaildtolist struct {

    // 仓code
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // LBX订单号
    
    UnitBizCode   string `json:"unit_biz_code,omitempty" xml:"unit_biz_code,omitempty"`
    

    // 1-出库单；2-入库单
    
    UnitType   string `json:"unit_type,omitempty" xml:"unit_type,omitempty"`
    

    // 出库lbx-下发仓、仓接单、部分出、全出 	 * 入库lbx-下发仓、仓接单、部分入、全入
    
    FulfilUniBizStatus   string `json:"fulfil_uni_biz_status,omitempty" xml:"fulfil_uni_biz_status,omitempty"`
    

    // 品基本信息
    
    TransferUnitOrderItemList   []Transferunitorderitemdtos `json:"transfer_unit_order_item_list,omitempty" xml:"transfer_unit_order_item_list,omitempty"`
    

}
