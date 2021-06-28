package waybill

// IsvResourceDo 
/* model for simplify = false
type IsvResourceDo struct {

    // 资源内容（当资源类型为TEMPLATE时，为空）
    
    ResourceContent   string `json:"resource_content,omitempty"`
    

    // 资源id
    
    ResourceId   int64 `json:"resource_id,omitempty"`
    

    // 资源名称
    
    ResourceName   string `json:"resource_name,omitempty"`
    

    // 资源类型
    
    ResourceType   string `json:"resource_type,omitempty"`
    

    // 资源url（当资源类型为打印项时，为空）
    
    ResourceUrl   string `json:"resource_url,omitempty"`
    

}
*/

// IsvResourceDo 
type IsvResourceDo struct {

    // 资源内容（当资源类型为TEMPLATE时，为空）
    ResourceContent   string `json:"resource_content,omitempty"`

    // 资源id
    ResourceId   int64 `json:"resource_id,omitempty"`

    // 资源名称
    ResourceName   string `json:"resource_name,omitempty"`

    // 资源类型
    ResourceType   string `json:"resource_type,omitempty"`

    // 资源url（当资源类型为打印项时，为空）
    ResourceUrl   string `json:"resource_url,omitempty"`

}
