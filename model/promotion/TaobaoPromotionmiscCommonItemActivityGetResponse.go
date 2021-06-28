package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动 APIResponse
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscCommonItemActivityGetResponse `json:"promotionmisc_common_item_activity_get_response,omitempty"` 
    TaobaoPromotionmiscCommonItemActivityGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscCommonItemActivityGetResponse struct {

    // 是否查询成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 优惠活动
    
    Activity  *struct {
        CommonItemActivity  *CommonItemActivity `json:"common_item_activity,omitempty"`
    } `json:"activity,omitempty"`
    

}
*/

type TaobaoPromotionmiscCommonItemActivityGetResponse struct {

    // 是否查询成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 优惠活动
    Activity   *CommonItemActivity `json:"activity,omitempty"`

}
