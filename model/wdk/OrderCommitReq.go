package wdk

// OrderCommitReq 
type OrderCommitReq struct {

    // 商品列表
    ItemConfirmInfos   []ItemConfirmInfo `json:"item_confirm_infos,omitempty"`

    // 订单号
    ExternalOrderNo   string `json:"external_order_no,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

}
