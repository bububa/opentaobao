package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠详情列表 APIResponse
taobao.promotionmisc.common.item.detail.list.get

查询通用单品优惠详情列表。
*/
type TaobaoPromotionmiscCommonItemDetailListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscCommonItemDetailListGetResponse `json:"promotionmisc_common_item_detail_list_get_response,omitempty"` 
    TaobaoPromotionmiscCommonItemDetailListGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscCommonItemDetailListGetResponse struct {

    // 是否查询成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 数据总数量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 活动详情列表
    
    DetailList  struct {
        CommonItemDetail  []CommonItemDetail `json:"common_item_detail,omitempty"`
    } `json:"detail_list,omitempty"`
    

}
*/

type TaobaoPromotionmiscCommonItemDetailListGetResponse struct {

    // 是否查询成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 数据总数量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 活动详情列表
    DetailList   []CommonItemDetail `json:"detail_list,omitempty"`

}
