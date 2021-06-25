package alsc

// VoucherTemplateOpenReq 
type VoucherTemplateOpenReq struct {

    // 品牌ID
    BrandId   string `json:"brand_id,omitempty"`

    // 已删除数据
    Deleted   bool `json:"deleted,omitempty"`

    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 第几页,从1开始计数
    PageNo   int64 `json:"page_no,omitempty"`

    // 第几页,从1开始计数
    PageSize   int64 `json:"page_size,omitempty"`

    // SaaS门店ID
    ShopId   string `json:"shop_id,omitempty"`

    // UNUSED,USED,NO_INVENTORY,INVALID,  未使用,使用中,无库存,已失效
    StatusList   []String `json:"status_list,omitempty"`

    // 最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 最后ID
    LastId   string `json:"last_id,omitempty"`

    // 模版ID
    VoucherTemplateId   string `json:"voucher_template_id,omitempty"`

}
