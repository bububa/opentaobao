package rhino

// ClothingInboundConfirmDto 
type ClothingInboundConfirmDto struct {

    // 入库数量
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 工厂ID，对应WMS002：Company
    
    FactoryId   int64 `json:"factory_id,omitempty" xml:"factory_id,omitempty"`
    

    // 入库单号，对应WMS002：ReceiptId
    
    InboundId   int64 `json:"inbound_id,omitempty" xml:"inbound_id,omitempty"`
    

    // 仓库ID，对应WMS002：WareHouse
    
    WarehouseId   int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
    

    // 入库完成时间
    
    InboundDate   string `json:"inbound_date,omitempty" xml:"inbound_date,omitempty"`
    

    // 物料描述，对应WMS002：ItemName
    
    MaterialDesc   string `json:"material_desc,omitempty" xml:"material_desc,omitempty"`
    

    // 物料ID，对应WMS002：Item
    
    MaterialId   int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
    

    // 电商平台订单号orderCode
    
    EcOrderCode   string `json:"ec_order_code,omitempty" xml:"ec_order_code,omitempty"`
    

    // sku 信息
    
    SkuBO   *SkuBo `json:"sku_b_o,omitempty" xml:"sku_b_o,omitempty"`
    

    // 入库单类型
    
    InboundType   int64 `json:"inbound_type,omitempty" xml:"inbound_type,omitempty"`
    

    // 生产订单号
    
    MesOrderId   string `json:"mes_order_id,omitempty" xml:"mes_order_id,omitempty"`
    

    // 成衣rfid编码
    
    Rfid   string `json:"rfid,omitempty" xml:"rfid,omitempty"`
    

    // toC订单Id
    
    DesiredId   string `json:"desired_id,omitempty" xml:"desired_id,omitempty"`
    

}
