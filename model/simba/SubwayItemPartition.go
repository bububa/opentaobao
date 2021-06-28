package simba

// SubwayItemPartition 
type SubwayItemPartition struct {

    // 每页数据大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 返回第几页的数据从1开始
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
    
    TotalItem   int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
    

    // 排序字段
    
    OrderField   string `json:"order_field,omitempty" xml:"order_field,omitempty"`
    

    // 排序，<br/>True:升级False:降序
    
    OrderBy   bool `json:"order_by,omitempty" xml:"order_by,omitempty"`
    

    // 商品列表
    
    ItemList   []SubwayItem `json:"item_list,omitempty" xml:"item_list,omitempty"`
    

}
