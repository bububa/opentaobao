package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.get

查询满就送活动
*/
type TaobaoPromotionmiscMjsActivityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_mjs_activity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 满就送活动信息。
    
    MjsPromotion   *MjsPromotion `json:"mjs_promotion,omitempty" xml:"