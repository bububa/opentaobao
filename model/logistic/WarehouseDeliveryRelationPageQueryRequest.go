package logistic

// WarehouseDeliveryRelationPageQueryRequest 
type WarehouseDeliveryRelationPageQueryRequest struct {

    // 网格仓外部编码 和fromResourceCode 2选1
    
    FromOrgResourceCode   string `json:"from_org_resource_code,omitempty" xml:"from_org_resource_code,omitempty"`
    

    // 网格仓外部编码列表
    
    FromOrgResourceCodeList   []string `json:"from_org_resource_code_list,omitempty" xml:"from_org_resource_code_list>string,omitempty"`
    

    // 网格仓资源编码 和fromOrgResourceCode 2选1
    
    FromResourceCode   string `json:"from_resource_code,omitempty" xml:"from_resource_code,omitempty"`
    

    // 网格仓资源编码列表
    
    FromResourceCodeList   []string `json:"from_resource_code_list,omitempty" xml:"from_resource_code_list>string,omitempty"`
    

    // 仓资源类型
    
    FromResourceType   string `json:"from_resource_type,omitempty" xml:"from_resource_type,omitempty"`
    

    // 网络编码
    
    NetworkCode   string `json:"network_code,omitempty" xml:"network_code,omitempty"`
    

    // 页码
    
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    

    // 分页大小，最大300
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 是否返回汇总数
    
    ShowTotal   int64 `json:"show_total,omitempty" xml:"show_total,omitempty"`
    

    // 自提点编码 和toResourceCode 2选1
    
    ToOrgResourceCode   string `json:"to_org_resource_code,omitempty" xml:"to_org_resource_code,omitempty"`
    

    // 自提点编码 和toOrgResourceCode 2选1
    
    ToResourceCode   string `json:"to_resource_code,omitempty" xml:"to_resource_code,omitempty"`
    

    // 自提点列表
    
    ToResourceCodeList   []string `json:"to_resource_code_list,omitempty" xml:"to_resource_code_list>string,omitempty"`
    

    // 自提点资源类型
    
    ToResourceType   string `json:"to_resource_type,omitempty" xml:"to_resource_type,omitempty"`
    

    // 开始时间
    
    StartGmtModified   string `json:"start_gmt_modified,omitempty" xml:"start_gmt_modified,omitempty"`
    

    // 结束时间
    
    EndGmtModified   string `json:"end_gmt_modified,omitempty" xml:"end_gmt_modified,omitempty"`
    

}
