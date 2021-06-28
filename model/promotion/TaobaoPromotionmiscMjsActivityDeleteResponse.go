package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
type TaobaoPromotionmiscMjsActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_mjs_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功删除活动。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"