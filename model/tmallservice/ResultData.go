package tmallservice

// Resultdata 
/* model for simplify = false
type Resultdata struct {

    // 创建时间
    
    CreateTime   string `json:"create_time,omitempty"`
    

    // 服务产品title
    
    Title   string `json:"title,omitempty"`
    

    // 服务产品介绍
    
    ServiceProductContent   string `json:"service_product_content,omitempty"`
    

    // 服务产品id
    
    ServiceProductId   int64 `json:"service_product_id,omitempty"`
    

    // 服务产品sku列表
    
    SimpleServiceSkuList  struct {
        SimpleServiceSkuDtOs  []SimpleServiceSkuDtOs `json:"simple_service_sku_dt_os,omitempty"`
    } `json:"simple_service_sku_list,omitempty"`
    

    // 服务名称
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 服务产品状态
    
    ServiceProductStatus   int64 `json:"service_product_status,omitempty"`
    

    // 服务产品类型
    
    ServiceProductType   int64 `json:"service_product_type,omitempty"`
    

}
*/

// Resultdata 
type Resultdata struct {

    // 创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 服务产品title
    Title   string `json:"title,omitempty"`

    // 服务产品介绍
    ServiceProductContent   string `json:"service_product_content,omitempty"`

    // 服务产品id
    ServiceProductId   int64 `json:"service_product_id,omitempty"`

    // 服务产品sku列表
    SimpleServiceSkuList   []SimpleServiceSkuDtOs `json:"simple_service_sku_list,omitempty"`

    // 服务名称
    ServiceCode   string `json:"service_code,omitempty"`

    // 服务产品状态
    ServiceProductStatus   int64 `json:"service_product_status,omitempty"`

    // 服务产品类型
    ServiceProductType   int64 `json:"service_product_type,omitempty"`

}
