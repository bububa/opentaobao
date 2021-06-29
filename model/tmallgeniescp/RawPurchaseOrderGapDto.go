package tmallgeniescp

// RawPurchaseOrderGapDto 
type RawPurchaseOrderGapDto struct {

    // 扩展参数
    
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    

    // 租户
    
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    

    // 关键日期值
    
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    

    // 二级物料-物料编码
    
    MaterielCode   string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
    

    // 二级物料-分仓目的地code
    
    LocationCodeTo   string `json:"location_code_to,omitempty" xml:"location_code_to,omitempty"`
    

    // 二级物料-供应商code
    
    LocationCode   string `json:"location_code,omitempty" xml:"location_code,omitempty"`
    

    // 二级物料PO GAP数
    
    RawPoGapQty   int64 `json:"raw_po_gap_qty,omitempty" xml:"raw_po_gap_qty,omitempty"`
    

}
