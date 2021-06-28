package alsc

// PullCardTemplateOpenReq 
/* model for simplify = false
type PullCardTemplateOpenReq struct {

    // 品牌id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 是否包含逻辑删除,数据下行时必须传递
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 更新时间,数据下行时必须传递
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 上次拉取的业务ID,数据下行时必须传递
    
    LastMaxId   string `json:"last_max_id,omitempty"`
    

    // 是否需要总数,分页查询卡模板时必须传递
    
    NeedCount   bool `json:"need_count,omitempty"`
    

    // 单次返回数量,数据下行时必须传递
    
    Num   int64 `json:"num,omitempty"`
    

    // 外部品牌id,与brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 外部门店id
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 分页参数,当前页码,分页查询卡模板时必须传递
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 分页参数,页面大小,分页查询卡模板时必须传递
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 是否查询详情,若需要详情,传true
    
    DetailRequired   bool `json:"detail_required,omitempty"`
    

}
*/

// PullCardTemplateOpenReq 
type PullCardTemplateOpenReq struct {

    // 品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 卡类型
    CardType   string `json:"card_type,omitempty"`

    // 是否包含逻辑删除,数据下行时必须传递
    Deleted   bool `json:"deleted,omitempty"`

    // 更新时间,数据下行时必须传递
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 上次拉取的业务ID,数据下行时必须传递
    LastMaxId   string `json:"last_max_id,omitempty"`

    // 是否需要总数,分页查询卡模板时必须传递
    NeedCount   bool `json:"need_count,omitempty"`

    // 单次返回数量,数据下行时必须传递
    Num   int64 `json:"num,omitempty"`

    // 外部品牌id,与brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部门店id
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 分页参数,当前页码,分页查询卡模板时必须传递
    PageNo   int64 `json:"page_no,omitempty"`

    // 分页参数,页面大小,分页查询卡模板时必须传递
    PageSize   int64 `json:"page_size,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 是否查询详情,若需要详情,传true
    DetailRequired   bool `json:"detail_required,omitempty"`

}
