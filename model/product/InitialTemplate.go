package product

// InitialTemplate 
type InitialTemplate struct {

    // 运费模板类型，可选值：freeshipping(全球免邮)，not(全球不发货)
    
    TemplateType   string `json:"template_type,omitempty" xml:"template_type,omitempty"`
    

    // 运费模板发货地，可选值：US（美国）,UK(英国),DE(德国),ES(西班牙),CN(中国)
    
    DispatchLocations   []string `json:"dispatch_locations,omitempty" xml:"dispatch_locations>string,omitempty"`
    

}
