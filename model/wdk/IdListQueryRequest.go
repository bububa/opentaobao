package wdk

// IdListQueryRequest 
/* model for simplify = false
type IdListQueryRequest struct {

    // 中台订单号
    
    BizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_id_list,omitempty"`
    

    // 淘系订单号
    
    TbBizIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tb_biz_id_list,omitempty"`
    

    // 渠道来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 经营店id
    
    StoreId   string `json:"store_id,omitempty"`
    

}
*/

// IdListQueryRequest 
type IdListQueryRequest struct {

    // 中台订单号
    BizIdList   []int64 `json:"biz_id_list,omitempty"`

    // 淘系订单号
    TbBizIdList   []int64 `json:"tb_biz_id_list,omitempty"`

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

}
