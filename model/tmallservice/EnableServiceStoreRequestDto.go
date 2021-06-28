package tmallservice

// EnableServiceStoreRequestDto 
/* model for simplify = false
type EnableServiceStoreRequestDto struct {

    // 门店名称
    
    StoreName   string `json:"store_name,omitempty"`
    

    // 是否启用
    
    Enable   bool `json:"enable,omitempty"`
    

    // 服务编码
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 服务sku列表
    
    ServiceSkuList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"service_sku_list,omitempty"`
    

    // 门店id
    
    StoreId   int64 `json:"store_id,omitempty"`
    

}
*/

// EnableServiceStoreRequestDto 
type EnableServiceStoreRequestDto struct {

    // 门店名称
    StoreName   string `json:"store_name,omitempty"`

    // 是否启用
    Enable   bool `json:"enable,omitempty"`

    // 服务编码
    ServiceCode   string `json:"service_code,omitempty"`

    // 服务sku列表
    ServiceSkuList   []string `json:"service_sku_list,omitempty"`

    // 门店id
    StoreId   int64 `json:"store_id,omitempty"`

}
