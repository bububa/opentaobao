package wdk

// IdListQueryRequest 
type IdListQueryRequest struct {

    // 中台订单号
    BizIdList   []Number `json:"biz_id_list,omitempty"`

    // 淘系订单号
    TbBizIdList   []Number `json:"tb_biz_id_list,omitempty"`

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

}
