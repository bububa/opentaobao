package wdk

// BomProcessDTO 
/* model for simplify = false
type BomProcessDTO struct {

    // 加工日期
    
    OccurrenceDate   string `json:"occurrence_date,omitempty"`
    

    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 唯一识别码
    
    Uuid   string `json:"uuid,omitempty"`
    

    // 部门编码
    
    DeptCode   string `json:"dept_code,omitempty"`
    

    // 加工类型
    
    OccurrenceType   string `json:"occurrence_type,omitempty"`
    

    // 单据编码
    
    BomProcessCode   string `json:"bom_process_code,omitempty"`
    

    // productItemInfos
    
    ProductItemInfos  struct {
        BomItemInfos  []BomItemInfos `json:"bom_item_infos,omitempty"`
    } `json:"product_item_infos,omitempty"`
    

    // materialItemInfos
    
    MaterialItemInfos  struct {
        BomItemInfos  []BomItemInfos `json:"bom_item_infos,omitempty"`
    } `json:"material_item_infos,omitempty"`
    

}
*/

// BomProcessDTO 
type BomProcessDTO struct {

    // 加工日期
    OccurrenceDate   string `json:"occurrence_date,omitempty"`

    // 店仓code，指的是库调对象，对应一个物理店或仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 唯一识别码
    Uuid   string `json:"uuid,omitempty"`

    // 部门编码
    DeptCode   string `json:"dept_code,omitempty"`

    // 加工类型
    OccurrenceType   string `json:"occurrence_type,omitempty"`

    // 单据编码
    BomProcessCode   string `json:"bom_process_code,omitempty"`

    // productItemInfos
    ProductItemInfos   []BomItemInfos `json:"product_item_infos,omitempty"`

    // materialItemInfos
    MaterialItemInfos   []BomItemInfos `json:"material_item_infos,omitempty"`

}
