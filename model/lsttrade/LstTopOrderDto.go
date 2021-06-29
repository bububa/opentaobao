package lsttrade

// LstTopOrderDto 
type LstTopOrderDto struct {
    // 实际支付金额
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
    // 订单创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 订单状态,等待买家付款   "waitbuyerpay";       *      * 等待卖家发货  "waitsellersend";       *      * 等待买家收货  "waitbuyerreceive";       *      * 交易成功   "success";       *      * 交易取消    "cancel";       *      * 已收货未结算  "confirm_goods_but_not_fund";       *      * [1688新新交易4.0] 已收货   "confirm_goods";
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 单价
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 买家id
    BuyerLoginId   string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
    // 商品条形码
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 产品名称
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    // 买家店铺名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 拍档id
    PartnerId   int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
    // 供应商名称
    SellerName   string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
    // 主订单id
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    // 子订单id
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    // 卖家loginId
    SellerLoginId   string `json:"seller_login_id,omitempty" xml:"seller_login_id,omitempty"`
    // 子订单支付时间
    GmtPayment   string `json:"gmt_payment,omitempty" xml:"gmt_payment,omitempty"`
    // 子订单发货时间
    GmtGoodsSend   string `json:"gmt_goods_send,omitempty" xml:"gmt_goods_send,omitempty"`
    // 子订单确认收货时间
    GmtGoodsReceived   string `json:"gmt_goods_received,omitempty" xml:"gmt_goods_received,omitempty"`
    // 子订单交易完结时间
    GmtCompleted   string `json:"gmt_completed,omitempty" xml:"gmt_completed,omitempty"`
    // 商家ERP编码ID
    OuterOfferId   string `json:"outer_offer_id,omitempty" xml:"outer_offer_id,omitempty"`
    // 订单上的收货地址
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    // true表示赠品
    Gift   bool `json:"gift,omitempty" xml:"gift,omitempty"`
    // 商家私域门店id
    OuterStoreId   string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
    // 商家业代id
    OuterOperatorId   string `json:"outer_operator_id,omitempty" xml:"outer_operator_id,omitempty"`
    // 订单状态,等待买家付款   "waitbuyerpay";       *      * 等待卖家发货  "waitsellersend";       *      * 等待买家收货  "waitbuyerreceive";       *      * 交易成功   "success";       *      * 交易取消    "cancel";       *      * 已收货未结算  "confirm_goods_but_not_fund";       *      * [1688新新交易4.0] 已收货   "confirm_goods";
    MainOrderStatus   string `json:"main_order_status,omitempty" xml:"main_order_status,omitempty"`
    // 仓库code
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 仓库名称
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    // 小店唯一ID
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 零售通业代id
    LstYdId   string `json:"lst_yd_id,omitempty" xml:"lst_yd_id,omitempty"`
    // 赠品订单对应的活动订单id，多个使用英文逗号','分割
    ActSubOrderIds   string `json:"act_sub_order_ids,omitempty" xml:"act_sub_order_ids,omitempty"`
    // 批次单
    GroupId   string `json:"group_id,omitempty" xml:"group_id,omitempty"`
    // 送达时间
    SignTime   string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
    // 是否组合产品
    CombineItem   bool `json:"combine_item,omitempty" xml:"combine_item,omitempty"`
}
