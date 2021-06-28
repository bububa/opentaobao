package product

// ItemTemplate 
type ItemTemplate struct {

    // 宝贝模板的id
    
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
    

    // 宝贝详情模板的名称
    
    TemplateName   string `json:"template_name,omitempty" xml:"template_name,omitempty"`
    

    // 用于区分宝贝模板属于内店和外店
    
    ShopType   int64 `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
    

}
