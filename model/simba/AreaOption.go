package simba

// AreaOption 
/* model for simplify = false
type AreaOption struct {

    // 地域id
    
    AreaId   int64 `json:"area_id,omitempty"`
    

    // 父地域id，若该字段为0表示该行政区为顶层，例如像北京，国外等。
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

    // 地域名称
    
    Name   string `json:"name,omitempty"`
    

    // 地域级别，目前自治区、省、直辖市是1，其他城市、地区是2
    
    Level   int64 `json:"level,omitempty"`
    

}
*/

// AreaOption 
type AreaOption struct {

    // 地域id
    AreaId   int64 `json:"area_id,omitempty"`

    // 父地域id，若该字段为0表示该行政区为顶层，例如像北京，国外等。
    ParentId   int64 `json:"parent_id,omitempty"`

    // 地域名称
    Name   string `json:"name,omitempty"`

    // 地域级别，目前自治区、省、直辖市是1，其他城市、地区是2
    Level   int64 `json:"level,omitempty"`

}
