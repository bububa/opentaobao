package viapi

// Detail 
type Detail struct {

    // 文本命中风险的分类
    
    Label   string `json:"label,omitempty" xml:"label,omitempty"`
    

    // 命中该风险的上下文信息
    
    Contexts   []Context `json:"contexts,omitempty" xml:"contexts,omitempty"`
    

}
