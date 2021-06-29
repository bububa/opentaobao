package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关联的活动权益 APIResponse
taobao.promotion.benefit.activity.update

修改卖家活动中关联的对应的权益。
*/
type TaobaoPromotionBenefitActivityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivityUpdateResponse
}

type TaobaoPromotionBenefitActivityUpdateResponse struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
