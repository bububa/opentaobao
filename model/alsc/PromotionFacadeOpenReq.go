package alsc

// PromotionFacadeOpenReq 
/* model for simplify = false
type PromotionFacadeOpenReq struct {

    // 品牌ID
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 1
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 外部品牌ID
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 第几页,从1开始计数
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 每页大小，默认20
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // STATUS_NEW,STATUS_WORKING,STATUS_SUSPEND,STATUS_END，STATUS_EXPIRED,未投放,执行中,已暂停,已终止,已终止
    
    StatusList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"status_list,omitempty"`
    

    // 最后修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 最后时间
    
    LastId   string `json:"last_id,omitempty"`
    

    // 促销ID
    
    PromotionId   string `json:"promotion_id,omitempty"`
    

}
*/

// PromotionFacadeOpenReq 
type PromotionFacadeOpenReq struct {

    // 品牌ID
    BrandId   string `json:"brand_id,omitempty"`

    // 1
    Deleted   bool `json:"deleted,omitempty"`

    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 第几页,从1开始计数
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页大小，默认20
    PageSize   int64 `json:"page_size,omitempty"`

    // STATUS_NEW,STATUS_WORKING,STATUS_SUSPEND,STATUS_END，STATUS_EXPIRED,未投放,执行中,已暂停,已终止,已终止
    StatusList   []string `json:"status_list,omitempty"`

    // 最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 最后时间
    LastId   string `json:"last_id,omitempty"`

    // 促销ID
    PromotionId   string `json:"promotion_id,omitempty"`

}
