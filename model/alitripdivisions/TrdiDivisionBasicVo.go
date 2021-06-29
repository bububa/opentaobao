package alitripdivisions

// TrdiDivisionBasicVo 
type TrdiDivisionBasicVo struct {

    // 国家码
    
    CountryCode2   string `json:"country_code2,omitempty" xml:"country_code2,omitempty"`
    

    // 国家名称
    
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    

    // 行政区划id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 纬度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 省
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 经度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 简称
    
    NameAbbr   string `json:"name_abbr,omitempty" xml:"name_abbr,omitempty"`
    

    // 英文名称
    
    NameEn   string `json:"name_en,omitempty" xml:"name_en,omitempty"`
    

    // 父节点id
    
    ParentId   int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    

    // 拼音
    
    Pinyin   string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
    

    // 简拼
    
    Py   string `json:"py,omitempty" xml:"py,omitempty"`
    

    // 层级树id
    
    TreeId   string `json:"tree_id,omitempty" xml:"tree_id,omitempty"`
    

    // 层级树名称
    
    TreeName   string `json:"tree_name,omitempty" xml:"tree_name,omitempty"`
    

}
