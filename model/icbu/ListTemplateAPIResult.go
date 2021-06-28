package icbu

// ListTemplateAPIResult 
/* model for simplify = false
type ListTemplateAPIResult struct {

    // 运费模板总数
    
    Total   int64 `json:"total,omitempty"`
    

    // 运费模板集合
    
    Items  struct {
        ShippinglineTemplate  []ShippinglineTemplate `json:"shippingline_template,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// ListTemplateAPIResult 
type ListTemplateAPIResult struct {

    // 运费模板总数
    Total   int64 `json:"total,omitempty"`

    // 运费模板集合
    Items   []ShippinglineTemplate `json:"items,omitempty"`

}
