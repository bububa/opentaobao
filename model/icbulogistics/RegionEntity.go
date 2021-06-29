package icbulogistics

// RegionEntity 
type RegionEntity struct {

    // 节点ID
    
    AreaId   int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
    

    // 节点名称拼音
    
    Pinyin   string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
    

    // 上级节点ID
    
    ParentId   int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    

    // 中文名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 层级
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 上级节点名称
    
    ParentName   string `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
    

    // 邮编
    
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 子节点列表
    
    Childrens   []Division `json:"childrens,omitempty" xml:"childrens,omitempty"`
    

    // 地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

}
