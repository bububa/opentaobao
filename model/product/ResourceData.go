package product

// ResourceData 
/* model for simplify = false
type ResourceData struct {

    // 资源列表
    
    Datas  struct {
        ResourceDataRecord  []ResourceDataRecord `json:"resource_data_record,omitempty"`
    } `json:"datas,omitempty"`
    

    // 资源名称
    
    Name   string `json:"name,omitempty"`
    

    // 资源类型
    
    Type   int64 `json:"type,omitempty"`
    

}
*/

// ResourceData 
type ResourceData struct {

    // 资源列表
    Datas   []ResourceDataRecord `json:"datas,omitempty"`

    // 资源名称
    Name   string `json:"name,omitempty"`

    // 资源类型
    Type   int64 `json:"type,omitempty"`

}
