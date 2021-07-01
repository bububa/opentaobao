package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除满就送活动 API返回值 
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
type TaobaoPromotionmiscMjsActivityDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel
}

// 删除满就送活动 成功返回结果
type TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"promotionmisc_mjs_activity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功删除活动。
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
