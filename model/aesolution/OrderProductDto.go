package aesolution

// OrderProductDto 
type OrderProductDto struct {

    // total product amount
    
    TotalProductAmount   *SimpleMoney `json:"total_product_amount,omitempty" xml:"total_product_amount,omitempty"`
    

    // child order status
    
    SonOrderStatus   string `json:"son_order_status,omitempty" xml:"son_order_status,omitempty"`
    

    // sku code
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // order show status
    
    ShowStatus   string `json:"show_status,omitempty" xml:"show_status,omitempty"`
    

    // Last delivery time
    
    SendGoodsTime   string `json:"send_goods_time,omitempty" xml:"send_goods_time,omitempty"`
    

    // Shipper type. "SELLER_SEND_GOODS": seller shipping; "WAREHOUSE_SEND_GOODS": warehouse delivery
    
    SendGoodsOperator   string `json:"send_goods_operator,omitempty" xml:"send_goods_operator,omitempty"`
    

    // product unit price
    
    ProductUnitPrice   *SimpleMoney `json:"product_unit_price,omitempty" xml:"product_unit_price,omitempty"`
    

    // product unit
    
    ProductUnit   string `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
    

    // product standard
    
    ProductStandard   string `json:"product_standard,omitempty" xml:"product_standard,omitempty"`
    

    // product snap Url
    
    ProductSnapUrl   string `json:"product_snap_url,omitempty" xml:"product_snap_url,omitempty"`
    

    // product name
    
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    

    // product main image url
    
    ProductImgUrl   string `json:"product_img_url,omitempty" xml:"product_img_url,omitempty"`
    

    // product id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // product count
    
    ProductCount   int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
    

    // order ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // fake one compensate three flag
    
    MoneyBack3x   bool `json:"money_back3x,omitempty" xml:"money_back3x,omitempty"`
    

    // buyer memo
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // logistics service name( key)
    
    LogisticsType   string `json:"logistics_type,omitempty" xml:"logistics_type,omitempty"`
    

    // logistics service show name
    
    LogisticsServiceName   string `json:"logistics_service_name,omitempty" xml:"logistics_service_name,omitempty"`
    

    // Logistics amount(sub-orders have no shipping costs, please ignore)
    
    LogisticsAmount   *SimpleMoney `json:"logistics_amount,omitempty" xml:"logistics_amount,omitempty"`
    

    // issue status (NO_ISSUE; IN_ISSUE; END_ISSUE)
    
    IssueStatus   string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
    

    // issue mode
    
    IssueMode   string `json:"issue_mode,omitempty" xml:"issue_mode,omitempty"`
    

    // goods prepare days
    
    GoodsPrepareTime   int64 `json:"goods_prepare_time,omitempty" xml:"goods_prepare_time,omitempty"`
    

    // fund status (NOT_PAY; PAY_SUCCESS; WAIT_SELLER_CHECK)
    
    FundStatus   string `json:"fund_status,omitempty" xml:"fund_status,omitempty"`
    

    // Limited time
    
    FreightCommitDay   string `json:"freight_commit_day,omitempty" xml:"freight_commit_day,omitempty"`
    

    // escrow fee rate
    
    EscrowFeeRate   string `json:"escrow_fee_rate,omitempty" xml:"escrow_fee_rate,omitempty"`
    

    // delivery time
    
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    

    // child order id
    
    ChildId   int64 `json:"child_id,omitempty" xml:"child_id,omitempty"`
    

    // Whether child orders can submit disputes
    
    CanSubmitIssue   bool `json:"can_submit_issue,omitempty" xml:"can_submit_issue,omitempty"`
    

    // buyer last name
    
    BuyerSignerLastName   string `json:"buyer_signer_last_name,omitempty" xml:"buyer_signer_last_name,omitempty"`
    

    // buyer first name
    
    BuyerSignerFirstName   string `json:"buyer_signer_first_name,omitempty" xml:"buyer_signer_first_name,omitempty"`
    

    // afflicate fee rate
    
    AfflicateFeeRate   string `json:"afflicate_fee_rate,omitempty" xml:"afflicate_fee_rate,omitempty"`
    

}
