package alsc

// PromotionFacadeOpenReq 
type PromotionFacadeOpenReq struct {
    // 品牌ID
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 1
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 第几页,从1开始计数
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 每页大小，默认20
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // STATUS_NEW,STATUS_WORKING,STATUS_SUSPEND,STATUS_END，STATUS_EXPIRED,未投放,执行中,已暂停,已终止,已终止
    StatusList   []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
    // 最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 最后时间
    LastId   string `json:"last_id,omitempty" xml:"last_id,omitempty"`
    // 促销ID
    PromotionId   string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
}
