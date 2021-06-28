package gameact

// AwardVO 
type AwardVO struct {

    // 奖品id
    
    AwardId   int64 `json:"award_id,omitempty" xml:"award_id,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 活动id
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    

    // 1:集分宝 2：红包 3：彩票 4：宝点  5：淘金币
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 根据类型展示不同的名称
    
    TypeName   string `json:"type_name,omitempty" xml:"type_name,omitempty"`
    

    // 数量
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 发奖单位
    
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    

    // 单位为分
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 分组编码
    
    GroupCode   int64 `json:"group_code,omitempty" xml:"group_code,omitempty"`
    

    // 分组描述
    
    GroupDesc   string `json:"group_desc,omitempty" xml:"group_desc,omitempty"`
    

    // 简短描述
    
    SimpleDesc   string `json:"simple_desc,omitempty" xml:"simple_desc,omitempty"`
    

    // 完整描述
    
    FullDesc   string `json:"full_desc,omitempty" xml:"full_desc,omitempty"`
    

    // 超链文案
    
    HyperlinkDesc   string `json:"hyperlink_desc,omitempty" xml:"hyperlink_desc,omitempty"`
    

    // 超链url
    
    HyperlinkUrl   string `json:"hyperlink_url,omitempty" xml:"hyperlink_url,omitempty"`
    

    // 加密串，对外流水号
    
    SerialNumber   string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
    

}
