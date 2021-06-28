package qimen

// StoreProcessCreateRequest 
/* model for simplify = false
type StoreProcessCreateRequest struct {

    // 加工单编码
    
    ProcessOrderCode   string `json:"processOrderCode,omitempty"`
    

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 单据类型(CNJG=仓内加工作业单)
    
    OrderType   string `json:"orderType,omitempty"`
    

    // 加工单创建时间(YYYY-MM-DD HH:MM:SS)
    
    OrderCreateTime   string `json:"orderCreateTime,omitempty"`
    

    // 计划加工时间(YYYY-MM-DD HH:MM:SS)
    
    PlanTime   string `json:"planTime,omitempty"`
    

    // 加工类型(1:仓内组合加工 2:仓内组合拆分)
    
    ServiceType   string `json:"serviceType,omitempty"`
    

    // 成品计划数量
    
    PlanQty   int64 `json:"planQty,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 加工商品列表
    
    Materialitems  struct {
        MaterialItem  []MaterialItem `json:"material_item,omitempty"`
    } `json:"materialitems,omitempty"`
    

    // 商品列表
    
    Productitems  struct {
        ProductItem  []ProductItem `json:"product_item,omitempty"`
    } `json:"productitems,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// StoreProcessCreateRequest 
type StoreProcessCreateRequest struct {

    // 加工单编码
    ProcessOrderCode   string `json:"processOrderCode,omitempty"`

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 单据类型(CNJG=仓内加工作业单)
    OrderType   string `json:"orderType,omitempty"`

    // 加工单创建时间(YYYY-MM-DD HH:MM:SS)
    OrderCreateTime   string `json:"orderCreateTime,omitempty"`

    // 计划加工时间(YYYY-MM-DD HH:MM:SS)
    PlanTime   string `json:"planTime,omitempty"`

    // 加工类型(1:仓内组合加工 2:仓内组合拆分)
    ServiceType   string `json:"serviceType,omitempty"`

    // 成品计划数量
    PlanQty   int64 `json:"planQty,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 加工商品列表
    Materialitems   []MaterialItem `json:"materialitems,omitempty"`

    // 商品列表
    Productitems   []ProductItem `json:"productitems,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
