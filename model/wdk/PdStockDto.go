package wdk

// PdStockDto 
/* model for simplify = false
type PdStockDto struct {

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 盘点类型，可选值：1：大盘  2：周盘 ；
    
    PdType   int64 `json:"pd_type,omitempty"`
    

    // 单据号
    
    PdOrderCode   string `json:"pd_order_code,omitempty"`
    

    // 店仓code，指的是盘点对象，对应一个物理店或仓编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 唯一识别码
    
    Uuid   string `json:"uuid,omitempty"`
    

    // itemList
    
    ItemList  struct {
        PdStockDetailDto  []PdStockDetailDto `json:"pd_stock_detail_dto,omitempty"`
    } `json:"item_list,omitempty"`
    

}
*/

// PdStockDto 
type PdStockDto struct {

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 盘点类型，可选值：1：大盘  2：周盘 ；
    PdType   int64 `json:"pd_type,omitempty"`

    // 单据号
    PdOrderCode   string `json:"pd_order_code,omitempty"`

    // 店仓code，指的是盘点对象，对应一个物理店或仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 唯一识别码
    Uuid   string `json:"uuid,omitempty"`

    // itemList
    ItemList   []PdStockDetailDto `json:"item_list,omitempty"`

}
