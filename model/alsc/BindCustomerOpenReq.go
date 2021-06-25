package alsc

// BindCustomerOpenReq 
type BindCustomerOpenReq struct {

    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 物理卡ID
    PhysicalCardId   string `json:"physical_card_id,omitempty"`

    // 请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 卡ID
    CardId   string `json:"card_id,omitempty"`

    // 品牌ID，和外部品牌ID必传1
    BrandId   string `json:"brand_id,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 顾客ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 门店ID
    ShopId   string `json:"shop_id,omitempty"`

    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty"`

}
