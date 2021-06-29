package ott

// MetaChartDO 
type MetaChartDO struct {
    // 排行标题
    Titles   []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
    // 排行类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
