package ascpchannel

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryData 
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryData struct {

    // 货品id
    
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // 物流货主id
    
    SourceUserId   int64 `json:"source_user_id,omitempty" xml:"source_user_id,omitempty"`
    

    // 有效状态
    
    InvStatus   int64 `json:"inv_status,omitempty" xml:"inv_status,omitempty"`
    

    // 可售总量
    
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 占用量
    
    LockQuantity   string `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
    

    // 预留库存总量
    
    ReservedQuantity   string `json:"reserved_quantity,omitempty" xml:"reserved_quantity,omitempty"`
    

    // 预留库存占用量
    
    ReservedLockQuantity   string `json:"reserved_lock_quantity,omitempty" xml:"reserved_lock_quantity,omitempty"`
    

    // 协议ID
    
    PublishOrderNo   string `json:"publish_order_no,omitempty" xml:"publish_order_no,omitempty"`
    

    // 销售市场中的计划ID，考拉中发布单号=计划ID
    
    TradeFutureInvId   int64 `json:"trade_future_inv_id,omitempty" xml:"trade_future_inv_id,omitempty"`
    

}
