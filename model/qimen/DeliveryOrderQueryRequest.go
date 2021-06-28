package qimen

// DeliveryOrderQueryRequest 
/* model for simplify = false
type DeliveryOrderQueryRequest struct {

    // 货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 仓库编码
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 发库单号
    
    OrderCode   string `json:"orderCode,omitempty"`
    

    // 仓储系统发库单号
    
    OrderId   string `json:"orderId,omitempty"`
    

    // 交易单号
    
    OrderSourceCode   string `json:"orderSourceCode,omitempty"`
    

    // 当前页
    
    Page   int64 `json:"page,omitempty"`
    

    // 每页orderLine条数(最多100条)
    
    PageSize   int64 `json:"pageSize,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

}
*/

// DeliveryOrderQueryRequest 
type DeliveryOrderQueryRequest struct {

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 发库单号
    OrderCode   string `json:"orderCode,omitempty"`

    // 仓储系统发库单号
    OrderId   string `json:"orderId,omitempty"`

    // 交易单号
    OrderSourceCode   string `json:"orderSourceCode,omitempty"`

    // 当前页
    Page   int64 `json:"page,omitempty"`

    // 每页orderLine条数(最多100条)
    PageSize   int64 `json:"pageSize,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}
