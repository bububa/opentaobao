package alsc

// VoucherStatusChangeOpenReq 
/* model for simplify = false
type VoucherStatusChangeOpenReq struct {

    // 品牌ID
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 顾客ID
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 操作ID
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 操作名字
    
    OperatorName   string `json:"operator_name,omitempty"`
    

    // 外部品牌ID
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 外部orderID
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 外部门店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 请求ID
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 门店ID
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // POS,移动门店
    
    Source   string `json:"source,omitempty"`
    

    // 优惠券集合
    
    VoucherIdList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"voucher_id_list,omitempty"`
    

    // /**      * 核销，正常到核销      */     NORMAL_ISUSED,     /**      * 反核销，已使用到正常（补发一张新的优惠券）      */     ISUSED_NORMAL,
    
    VoucherStatusAction   string `json:"voucher_status_action,omitempty"`
    

    // 待核销点数，数组下标和券id对齐
    
    VoucherPointList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"voucher_point_list,omitempty"`
    

}
*/

// VoucherStatusChangeOpenReq 
type VoucherStatusChangeOpenReq struct {

    // 品牌ID
    BrandId   string `json:"brand_id,omitempty"`

    // 顾客ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 操作ID
    OperatorId   string `json:"operator_id,omitempty"`

    // 操作名字
    OperatorName   string `json:"operator_name,omitempty"`

    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部orderID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 门店ID
    ShopId   string `json:"shop_id,omitempty"`

    // POS,移动门店
    Source   string `json:"source,omitempty"`

    // 优惠券集合
    VoucherIdList   []string `json:"voucher_id_list,omitempty"`

    // /**      * 核销，正常到核销      */     NORMAL_ISUSED,     /**      * 反核销，已使用到正常（补发一张新的优惠券）      */     ISUSED_NORMAL,
    VoucherStatusAction   string `json:"voucher_status_action,omitempty"`

    // 待核销点数，数组下标和券id对齐
    VoucherPointList   []string `json:"voucher_point_list,omitempty"`

}
