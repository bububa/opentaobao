package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_item_activity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 单品优惠活动信息。
    
    ItemPromotion   *ItemPromotion `json:"item_promotion,omitempty" xml:"