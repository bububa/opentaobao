package wdk

// SupplierOrderQueryListRequest 
type SupplierOrderQueryListRequest struct {

    // 订单渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

    // 淘宝订单号
    TbBizIdList   []Number `json:"tb_biz_id_list,omitempty"`

    // 盒马订单号
    BizIdList   []Number `json:"biz_id_list,omitempty"`

}
