package qimen

// EntryOrderQueryRequest 
/* model for simplify = false
type EntryOrderQueryRequest struct {

    // 货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 仓库编码
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 入库单编码
    
    EntryOrderCode   string `json:"entryOrderCode,omitempty"`
    

    // 仓储系统入库单ID
    
    EntryOrderId   string `json:"entryOrderId,omitempty"`
    

    // 当前页(从1开始)
    
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

// EntryOrderQueryRequest 
type EntryOrderQueryRequest struct {

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 入库单编码
    EntryOrderCode   string `json:"entryOrderCode,omitempty"`

    // 仓储系统入库单ID
    EntryOrderId   string `json:"entryOrderId,omitempty"`

    // 当前页(从1开始)
    Page   int64 `json:"page,omitempty"`

    // 每页orderLine条数(最多100条)
    PageSize   int64 `json:"pageSize,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}
