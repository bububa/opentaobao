package qimen

// ItemLackReportRequest 
type ItemLackReportRequest struct {

    // 仓库编码
    
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    

    // ERP的发货单编码
    
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    

    // 仓库系统的发货单编码
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    

    // 缺货回告创建时间(YYYY-MM-DD HH:mm:ss)
    
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    

    // 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
    
    OutBizCode   string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
    

    // 缺货商品列表
    
    Items   []Item `json:"items,omitempty" xml:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *Map `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

}
