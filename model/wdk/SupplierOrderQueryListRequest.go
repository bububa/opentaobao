package wdk

// SupplierOrderQueryListRequest 
type SupplierOrderQueryListRequest struct {
    // 订单渠道来源
    OrderFrom   int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
    // 商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
    // 淘宝订单号
    TbBizIdList   []int64 `json:"tb_biz_id_list,omitempty" xml:"tb_biz_id_list>int64,omitempty"`
    // 盒马订单号
    BizIdList   []int64 `json:"biz_id_list,omitempty" xml:"biz_id_list>int64,omitempty"`
}
