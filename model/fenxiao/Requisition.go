package fenxiao

// Requisition 
/* model for simplify = false
type Requisition struct {

    // 合作申请ID
    
    RequisitionId   int64 `json:"requisition_id,omitempty"`
    

    // 分销商id
    
    DistributorId   int64 `json:"distributor_id,omitempty"`
    

    // 分销商nick
    
    DistributorNick   string `json:"distributor_nick,omitempty"`
    

    // 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
    
    Status   int64 `json:"status,omitempty"`
    

    // 是否消保(0-不是、1-是)
    
    DistIsXiaobao   int64 `json:"dist_is_xiaobao,omitempty"`
    

    // 店铺地址
    
    DistShopAddress   string `json:"dist_shop_address,omitempty"`
    

    // 店铺星级
    
    DistLevel   int64 `json:"dist_level,omitempty"`
    

    // 好评率
    
    DistAppraise   int64 `json:"dist_appraise,omitempty"`
    

    // 开店时间
    
    DistOpenDate   string `json:"dist_open_date,omitempty"`
    

    // 主营类目
    
    DistCategory   int64 `json:"dist_category,omitempty"`
    

    // 申请时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 分销申请加盟时，给供应商的留言
    
    DistMessage   string `json:"dist_message,omitempty"`
    

    // 主营类目名称
    
    DistCategoryName   string `json:"dist_category_name,omitempty"`
    

}
*/

// Requisition 
type Requisition struct {

    // 合作申请ID
    RequisitionId   int64 `json:"requisition_id,omitempty"`

    // 分销商id
    DistributorId   int64 `json:"distributor_id,omitempty"`

    // 分销商nick
    DistributorNick   string `json:"distributor_nick,omitempty"`

    // 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
    Status   int64 `json:"status,omitempty"`

    // 是否消保(0-不是、1-是)
    DistIsXiaobao   int64 `json:"dist_is_xiaobao,omitempty"`

    // 店铺地址
    DistShopAddress   string `json:"dist_shop_address,omitempty"`

    // 店铺星级
    DistLevel   int64 `json:"dist_level,omitempty"`

    // 好评率
    DistAppraise   int64 `json:"dist_appraise,omitempty"`

    // 开店时间
    DistOpenDate   string `json:"dist_open_date,omitempty"`

    // 主营类目
    DistCategory   int64 `json:"dist_category,omitempty"`

    // 申请时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 分销申请加盟时，给供应商的留言
    DistMessage   string `json:"dist_message,omitempty"`

    // 主营类目名称
    DistCategoryName   string `json:"dist_category_name,omitempty"`

}
