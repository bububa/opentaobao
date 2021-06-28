package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动列表 APIResponse
taobao.promotionmisc.mjs.activity.list.get

查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
*/
type TaobaoPromotionmiscMjsActivityListGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscMjsActivityListGetResponse
}

type TaobaoPromotionmiscMjsActivityListGetResponse struct {
    XMLName xml.Name `xml:"promotionmisc_mjs_activity_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只包含活动的主要信息，如activity_id，aame，description，start_time，end_time，type，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
    
    MjsPromotionList   []MjsPromotion `json:"mjs_promotion_list,omitempty" xml:"mjs_promotion_list>mjs_promotion,omitempty"`
    
    
    // 记录总条数。
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
