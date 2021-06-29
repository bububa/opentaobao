package moziacl

// PageRolePermissionResult 
type PageRolePermissionResult struct {

    // 角色下权限总数量
    
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
    

    // 角色下的权限列表数据
    
    Datas   []PermissionEntity `json:"datas,omitempty" xml:"datas,omitempty"`
    

    // 是否处理成功，成功则返回true
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 请求唯一id
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 要查询的角色name
    
    RoleName   string `json:"role_name,omitempty" xml:"role_name,omitempty"`
    

    // 每页数量
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 响应message，若失败则返回失败原因
    
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    

    // 扩展字段，与入参扩展字段值相同
    
    ResponseMetaData   string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
    

    // 响应code
    
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    

}
