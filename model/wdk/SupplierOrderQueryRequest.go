package wdk

// SupplierOrderQueryRequest 
type SupplierOrderQueryRequest struct {

    // 订单来源，4：盒马，3：饿了么
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

    // 实际售卖商家code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

    // 结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 分页页码，从0开始
    PageIndex   int64 `json:"page_index,omitempty"`

    // 分页大小，默认200
    PageSize   int64 `json:"page_size,omitempty"`

    // 下单终端: APP 、POS
    OrderClient   string `json:"order_client,omitempty"`

    // 订单状态: PAID / PACAKAGED / SUCCESS
    OrderStatus   []String `json:"order_status,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

}
