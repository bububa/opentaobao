package simba

// ADGroupPage 
/* model for simplify = false
type ADGroupPage struct {

    // 每页数据大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 返回第几页的数据从1开始
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
    
    TotalItem   int64 `json:"total_item,omitempty"`
    

    // 商品对象列表
    
    AdgroupList  struct {
        ADGroup  []ADGroup `json:"ad_group,omitempty"`
    } `json:"adgroup_list,omitempty"`
    

    // 模板规则
    
    Schedule   string `json:"schedule,omitempty"`
    

    // 模板名称
    
    Name   string `json:"name,omitempty"`
    

    // 模板id
    
    Id   int64 `json:"id,omitempty"`
    

}
*/

// ADGroupPage 
type ADGroupPage struct {

    // 每页数据大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 返回第几页的数据从1开始
    PageNo   int64 `json:"page_no,omitempty"`

    // 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
    TotalItem   int64 `json:"total_item,omitempty"`

    // 商品对象列表
    AdgroupList   []ADGroup `json:"adgroup_list,omitempty"`

    // 模板规则
    Schedule   string `json:"schedule,omitempty"`

    // 模板名称
    Name   string `json:"name,omitempty"`

    // 模板id
    Id   int64 `json:"id,omitempty"`

}
