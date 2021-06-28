package wlb

// WlbItemInventoryLog 
type WlbItemInventoryLog struct {

    // 库存变更ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 用户ID
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 库存操作作类型CHU_KU 1-出库RU_KU 2-入库FREEZE 3-冻结THAW 4-解冻CHECK_FREEZE 5-冻结确认CHANGE_KU 6-库存类型变更
    
    OpType   string `json:"op_type,omitempty" xml:"op_type,omitempty"`
    

    // 库存操作者ID
    
    OpUserId   int64 `json:"op_user_id,omitempty" xml:"op_user_id,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 批次号
    
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 订单号
    
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    

    // 订单商品ID
    
    OrderItemId   int64 `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
    

    // 处理数量变化值
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 创建日期
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 结果值
    
    ResultQuantity   int64 `json:"result_quantity,omitempty" xml:"result_quantity,omitempty"`
    

    // VENDIBLE  1-可销售;FREEZE  201-冻结库存;ONWAY  301-在途库存;DEFECT  101-残存品;ENGINE_DAMAGE 102-机损;BOX_DAMAGE 103-箱损
    
    InventType   string `json:"invent_type,omitempty" xml:"invent_type,omitempty"`
    

}
