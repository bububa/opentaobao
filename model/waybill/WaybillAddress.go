package waybill

// WaybillAddress 
/* model for simplify = false
type WaybillAddress struct {

    // 区名称（三级地址）
    
    Area   string `json:"area,omitempty"`
    

    // 一级地址（省、直辖市
    
    Province   string `json:"province,omitempty"`
    

    // 街道\镇名称（四级地址）
    
    Town   string `json:"town,omitempty"`
    

    // 详细地址
    
    AddressDetail   string `json:"address_detail,omitempty"`
    

    // 市名称（二级地址）
    
    City   string `json:"city,omitempty"`
    

    // 末级地址
    
    DivisionId   int64 `json:"division_id,omitempty"`
    

    // waybill 地址记录ID(非地址库ID)
    
    WaybillAddressId   int64 `json:"waybill_address_id,omitempty"`
    

}
*/

// WaybillAddress 
type WaybillAddress struct {

    // 区名称（三级地址）
    Area   string `json:"area,omitempty"`

    // 一级地址（省、直辖市
    Province   string `json:"province,omitempty"`

    // 街道\镇名称（四级地址）
    Town   string `json:"town,omitempty"`

    // 详细地址
    AddressDetail   string `json:"address_detail,omitempty"`

    // 市名称（二级地址）
    City   string `json:"city,omitempty"`

    // 末级地址
    DivisionId   int64 `json:"division_id,omitempty"`

    // waybill 地址记录ID(非地址库ID)
    WaybillAddressId   int64 `json:"waybill_address_id,omitempty"`

}
