package wdk

// SupplierOrderQueryListRequest 
/* model for simplify = false
type SupplierOrderQueryListRequest struct {

    // 订单渠道来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 商场code
    
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`
    

    // 淘宝订单号
    
    TbBizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tb_biz_id_list,omitempty"`
    

    // 盒马订单号
    
    BizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_id_list,omitempty"`
    

}
*/

// SupplierOrderQueryListRequest 
type SupplierOrderQueryListRequest struct {

    // 订单渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

    // 淘宝订单号
    TbBizIdList   []int64 `json:"tb_biz_id_list,omitempty"`

    // 盒马订单号
    BizIdList   []int64 `json:"biz_id_list,omitempty"`

}
