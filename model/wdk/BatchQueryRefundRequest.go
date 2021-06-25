package wdk

// BatchQueryRefundRequest 
type BatchQueryRefundRequest struct {

    // 结束时间，必填
    EndTime   string `json:"end_time,omitempty"`

    // 下单终端：APP/POS，非必填
    OrderClient   string `json:"order_client,omitempty"`

    // 页码，从0开始，必填
    PageIndex   int64 `json:"page_index,omitempty"`

    // 页大小，必填
    PageSize   int64 `json:"page_size,omitempty"`

    // 开始时间，必填
    StartTime   string `json:"start_time,omitempty"`

    // 店铺ID列表，order_from=4时必填，其他非必填
    StoreIds   []String `json:"store_ids,omitempty"`

    // 废弃字段
    SyncStatus   string `json:"sync_status,omitempty"`

    // 渠道来源 3：饿了么 4：盒马&淘鲜达 20：商家自有渠道，必填
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 渠道店id，order_from=4时非必填，其他必填
    ShopId   string `json:"shop_id,omitempty"`

    // 经营店id，非必填
    StoreId   string `json:"store_id,omitempty"`

    // aelophy 翱象商家必填
    BizType   string `json:"biz_type,omitempty"`

}
