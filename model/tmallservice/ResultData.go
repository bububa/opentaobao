package tmallservice

// Resultdata 
type Resultdata struct {

    // 创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 服务产品title
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 服务产品介绍
    
    ServiceProductContent   string `json:"service_product_content,omitempty" xml:"service_product_content,omitempty"`
    

    // 服务产品id
    
    ServiceProductId   int64 `json:"service_product_id,omitempty" xml:"service_product_id,omitempty"`
    

    // 服务产品sku列表
    
    SimpleServiceSkuList   []SimpleServiceSkuDtOs `json:"simple_service_sku_list,omitempty" xml:"simple_service_sku_list,omitempty"`
    

    // 服务名称
    
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    

    // 服务产品状态
    
    ServiceProductStatus   int64 `json:"service_product_status,omitempty" xml:"service_product_status,omitempty"`
    

    // 服务产品类型
    
    ServiceProductType   int64 `json:"service_product_type,omitempty" xml:"service_product_type,omitempty"`
    

}
