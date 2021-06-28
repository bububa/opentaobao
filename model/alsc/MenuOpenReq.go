package alsc

// MenuOpenReq 
type MenuOpenReq struct {

    // 品牌ID
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 已删除数据
    
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    

    // 是否需要总数
    
    NeedCount   bool `json:"need_count,omitempty" xml:"need_count,omitempty"`
    

    // 外部品牌ID
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

    // 外部门店ID
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 第几页,从1开始计数
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 每页大小，默认20
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // SaaS门店ID
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 最后修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 最后ID
    
    LastId   string `json:"last_id,omitempty" xml:"last_id,omitempty"`
    

    // 特价菜单ID
    
    MenuId   string `json:"menu_id,omitempty" xml:"menu_id,omitempty"`
    

}
