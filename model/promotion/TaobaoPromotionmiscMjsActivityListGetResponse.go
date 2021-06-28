package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动列表 APIResponse
taobao.promotionmisc.mjs.activity.list.get

查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
*/
type TaobaoPromotionmiscMjsActivityListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscMjsActivityListGetResponse `json:"promotionmisc_mjs_activity_list_get_response,omitempty"` 
    TaobaoPromotionmiscMjsActivityListGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscMjsActivityListGetResponse struct {

    // 只包含活动的主要信息，如activity_id，aame，description，start_time，end_time，type，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
    
    MjsPromotionList  struct {
        MjsPromotion  []MjsPromotion `json:"mjs_promotion,omitempty"`
    } `json:"mjs_promotion_list,omitempty"`
    

    // 记录总条数。
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoPromotionmiscMjsActivityListGetResponse struct {

    // 只包含活动的主要信息，如activity_id，aame，description，start_time，end_time，type，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
    MjsPromotionList   []MjsPromotion `json:"mjs_promotion_list,omitempty"`

    // 记录总条数。
    TotalCount   int64 `json:"total_count,omitempty"`

}
