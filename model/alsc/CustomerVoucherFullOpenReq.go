package alsc

// CustomerVoucherFullOpenReq 
type CustomerVoucherFullOpenReq struct {

    // 品牌ID
    BrandId   string `json:"brand_id,omitempty"`

    // 顾客ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 第几页,从1开始计数
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页大小，默认20
    PageSize   int64 `json:"page_size,omitempty"`

    // SaaS门店ID
    ShopId   string `json:"shop_id,omitempty"`

    // 优惠券状态 NORMAL，DELETED，ISUSED
    VoucherStatusList   []String `json:"voucher_status_list,omitempty"`

}
