package aesolution

// GlobalAeopTpChildOrderDto 
type GlobalAeopTpChildOrderDto struct {
    // snapshot ID
    SnapshotId   string `json:"snapshot_id,omitempty" xml:"snapshot_id,omitempty"`
    // How many products in one piece or lot
    LotNum   int64 `json:"lot_num,omitempty" xml:"lot_num,omitempty"`
    // logistics amount
    LogisticsAmount   *GlobalMoneyStr `json:"logistics_amount,omitempty" xml:"logistics_amount,omitempty"`
    // goods prepare days
    GoodsPrepareTime   int64 `json:"goods_prepare_time,omitempty" xml:"goods_prepare_time,omitempty"`
    // Extended attributes of product
    ProductAttributes   string `json:"product_attributes,omitempty" xml:"product_attributes,omitempty"`
    // buyer memo
    BuyerMemo   string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
    // refund info
    RefundInfo   *GlobalAeopTpRefundInfoDto `json:"refund_info,omitempty" xml:"refund_info,omitempty"`
    // product_unit
    ProductUnit   string `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
    // order ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // logistics type
    LogisticsType   string `json:"logistics_type,omitempty" xml:"logistics_type,omitempty"`
    // frozen status（NO_FROZEN; IN_FROZEN）
    FrozenStatus   string `json:"frozen_status,omitempty" xml:"frozen_status,omitempty"`
    // issue status
    IssueStatus   string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
    // order sort id
    OrderSortId   int64 `json:"order_sort_id,omitempty" xml:"order_sort_id,omitempty"`
    // affiliate fee rate
    AfflicateFeeRate   string `json:"afflicate_fee_rate,omitempty" xml:"afflicate_fee_rate,omitempty"`
    // order amount
    InitOrderAmt   *GlobalMoneyStr `json:"init_order_amt,omitempty" xml:"init_order_amt,omitempty"`
    // child issue info
    ChildIssueInfo   *GlobalAeopTpIssueInfoDto `json:"child_issue_info,omitempty" xml:"child_issue_info,omitempty"`
    // logistics service name
    LogisticsServiceName   string `json:"logistics_service_name,omitempty" xml:"logistics_service_name,omitempty"`
    // order loan info
    LoanInfo   *GlobalAeopTpLoanInfoDto `json:"loan_info,omitempty" xml:"loan_info,omitempty"`
    // product snapshot Url
    ProductSnapUrl   string `json:"product_snap_url,omitempty" xml:"product_snap_url,omitempty"`
    // Order Status：PLACE_ORDER_SUCCESS;  IN_CANCEL;  WAIT_SELLER_SEND_GOODS;  SELLER_PART_SEND_GOODS;  WAIT_BUYER_ACCEPT_GOODS;  FUND_PROCESSING; IN_ISSUE;  IN_FROZEN;  WAIT_SELLER_EXAMINE_MONEY;  RISK_CONTROL.
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // SKU code
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // SELLER_SEND_GOODS" or "WAREHOUSE_SEND_GOODS"
    SendGoodsOperator   string `json:"send_goods_operator,omitempty" xml:"send_goods_operator,omitempty"`
    // product ID
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // product quantity
    ProductCount   int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
    // fund status (NOT_PAY；PAY_SUCCESS)
    FundStatus   string `json:"fund_status,omitempty" xml:"fund_status,omitempty"`
    // escrow fee rate
    EscrowFeeRate   string `json:"escrow_fee_rate,omitempty" xml:"escrow_fee_rate,omitempty"`
    // product image Url
    ProductImgUrl   string `json:"product_img_url,omitempty" xml:"product_img_url,omitempty"`
    // child order ID
    ChildOrderId   string `json:"child_order_id,omitempty" xml:"child_order_id,omitempty"`
    // product price
    ProductPrice   *GlobalMoneyStr `json:"product_price,omitempty" xml:"product_price,omitempty"`
    // product name
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    // snapshot small photo path
    SnapshotSmallPhotoPath   string `json:"snapshot_small_photo_path,omitempty" xml:"snapshot_small_photo_path,omitempty"`
    // Product specifications
    ProductStandard   string `json:"product_standard,omitempty" xml:"product_standard,omitempty"`
    // cainiaoInternationalWarehouse
    LogisticsWarehouseType   string `json:"logistics_warehouse_type,omitempty" xml:"logistics_warehouse_type,omitempty"`
    // discount amount of child order
    ChildOrderDiscountInfo   *GlobalMoneyStr `json:"child_order_discount_info,omitempty" xml:"child_order_discount_info,omitempty"`
    // An amount to adjust the product price on, for example, if the seller wants to give buyer a personal discount by the "Adjust price"  button in order management backend.
    ChildOrderAdjustAmount   *GlobalMoneyStr `json:"child_order_adjust_amount,omitempty" xml:"child_order_adjust_amount,omitempty"`
    // discount detail list for each child order
    ChildOrderDiscountDetailList   []GlobalAeopTpSaleDiscountInfo `json:"child_order_discount_detail_list,omitempty" xml:"child_order_discount_detail_list>global_aeop_tp_sale_discount_info,omitempty"`
}
