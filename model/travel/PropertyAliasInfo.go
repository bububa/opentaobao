package travel

// PropertyAliasInfo 
type PropertyAliasInfo struct {

    // 销售属性的pid和vid，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid
    
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    

    // 属性具体别名值
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

}
