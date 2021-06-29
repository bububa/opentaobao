package koubeimall

// ItemServices 
type ItemServices struct {

    // 服务名称列表
    
    ServiceNameList   []string `json:"service_name_list,omitempty" xml:"service_name_list>string,omitempty"`
    

    // 服务明细，点击弹层显示
    
    ServicesDetailList   []ItemServicesDetail `json:"services_detail_list,omitempty" xml:"services_detail_list,omitempty"`
    

    // 服务列表
    
    ServiceCodeList   []string `json:"service_code_list,omitempty" xml:"service_code_list>string,omitempty"`
    

}
