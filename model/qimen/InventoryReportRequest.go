package qimen

// InventoryReportRequest 
/* model for simplify = false
type InventoryReportRequest struct {

    // 总页数
    
    TotalPage   int64 `json:"totalPage,omitempty"`
    

    // 当前页(从1开始)
    
    CurrentPage   int64 `json:"currentPage,omitempty"`
    

    // 每页记录的条数
    
    PageSize   int64 `json:"pageSize,omitempty"`
    

    // 仓库编码
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 盘点单编码
    
    CheckOrderCode   string `json:"checkOrderCode,omitempty"`
    

    // 仓储系统的盘点单编码
    
    CheckOrderId   string `json:"checkOrderId,omitempty"`
    

    // 货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 盘点时间(YYYY-MM-DD HH:MM:SS)
    
    CheckTime   string `json:"checkTime,omitempty"`
    

    // 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
    
    OutBizCode   string `json:"outBizCode,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 商品库存信息列表
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 变动类型：CHECK=盘点 ADJUST=调整
    
    AdjustType   string `json:"adjustType,omitempty"`
    

}
*/

// InventoryReportRequest 
type InventoryReportRequest struct {

    // 总页数
    TotalPage   int64 `json:"totalPage,omitempty"`

    // 当前页(从1开始)
    CurrentPage   int64 `json:"currentPage,omitempty"`

    // 每页记录的条数
    PageSize   int64 `json:"pageSize,omitempty"`

    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 盘点单编码
    CheckOrderCode   string `json:"checkOrderCode,omitempty"`

    // 仓储系统的盘点单编码
    CheckOrderId   string `json:"checkOrderId,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 盘点时间(YYYY-MM-DD HH:MM:SS)
    CheckTime   string `json:"checkTime,omitempty"`

    // 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
    OutBizCode   string `json:"outBizCode,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 商品库存信息列表
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 变动类型：CHECK=盘点 ADJUST=调整
    AdjustType   string `json:"adjustType,omitempty"`

}
