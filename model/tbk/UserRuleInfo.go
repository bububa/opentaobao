package tbk

// UserRuleInfo 
/* model for simplify = false
type UserRuleInfo struct {

    // 用户在TOP上的openId
    
    OpenId   string `json:"open_id,omitempty"`
    

    // 每条记录离线任务生成，代表当时离线任务的时间戳
    
    Version   int64 `json:"version,omitempty"`
    

    // 用户对应的商品详细信息
    
    ItemList  struct {
        TaobaoTbkCartCouponExpireUserQueryMapData  []TaobaoTbkCartCouponExpireUserQueryMapData `json:"taobao_tbk_cart_coupon_expire_user_query_map_data,omitempty"`
    } `json:"item_list,omitempty"`
    

}
*/

// UserRuleInfo 
type UserRuleInfo struct {

    // 用户在TOP上的openId
    OpenId   string `json:"open_id,omitempty"`

    // 每条记录离线任务生成，代表当时离线任务的时间戳
    Version   int64 `json:"version,omitempty"`

    // 用户对应的商品详细信息
    ItemList   []TaobaoTbkCartCouponExpireUserQueryMapData `json:"item_list,omitempty"`

}
