package logistic

// ResourceGroupPageQueryRequest 
type ResourceGroupPageQueryRequest struct {

    // 区块编码
    
    AreaCode   string `json:"area_code,omitempty" xml:"area_code,omitempty"`
    

    // 网格仓外部编码
    
    FromOrgResourceCode   string `json:"from_org_resource_code,omitempty" xml:"from_org_resource_code,omitempty"`
    

    // 网格仓编码
    
    FromResourceCode   string `json:"from_resource_code,omitempty" xml:"from_resource_code,omitempty"`
    

    // from资源类型
    
    FromResourceType   string `json:"from_resource_type,omitempty" xml:"from_resource_type,omitempty"`
    

    // 自提点编码
    
    GroupResourceCodeList   []string `json:"group_resource_code_list,omitempty" xml:"group_resource_code_list>string,omitempty"`
    

    // 网络名称
    
    NetworkCode   string `json:"network_code,omitempty" xml:"network_code,omitempty"`
    

    // 页码，1开始
    
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    

    // 页面大小，上限50
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
