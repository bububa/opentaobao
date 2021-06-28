package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动列表 APIResponse
taobao.promotionmisc.common.item.activity.list.get

查询通用单品优惠活动列表。
*/
type TaobaoPromotionmiscCommonItemActivityListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscCommonItemActivityListGetResponse `json:"promotionmisc_common_item_activity_list_get_response,omitempty"` 
    TaobaoPromotionmiscCommonItemActivityListGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscCommonItemActivityListGetResponse struct {

    // 是否查询成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 数据总数量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 营销活动列表
    
    ActivityList  struct {
        CommonItemActivity  []CommonItemActivity `json:"common_item_activity,omitempty"`
    } `json:"activity_list,omitempty"`
    

}
*/

type TaobaoPromotionmiscCommonItemActivityListGetResponse struct {

    // 是否查询成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 数据总数量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 营销活动列表
    ActivityList   []CommonItemActivity `json:"activity_list,omitempty"`

}
