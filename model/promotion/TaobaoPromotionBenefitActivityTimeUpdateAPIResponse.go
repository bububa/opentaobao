package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新关联活动有效时间 API返回值 
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间
*/
type TaobaoPromotionBenefitActivityTimeUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivityTimeUpdateAPIResponseModel
}

// 更新关联活动有效时间 成功返回结果
type TaobaoPromotionBenefitActivityTimeUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_time_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
