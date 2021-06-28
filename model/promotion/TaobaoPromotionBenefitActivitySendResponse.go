package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动权益发放接口 APIResponse
taobao.promotion.benefit.activity.send

活动权益发放接口，用于卖家针对活动进行权益发放
*/
type TaobaoPromotionBenefitActivitySendAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivitySendResponse
}

type TaobaoPromotionBenefitActivitySendResponse struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 事务id;发放时，不传入事务id,发放返回结果中会包含事务id；若发放失败，使用返回的事务id进行调用，则可以重试失败的操作
    
    SessionId   string `json:"session_id,omitempty" xml:"session_id,omitempty"`

    
    // 返回活动详情级别的权益发放情况
    
    ResultList   []BenefitSendResultExt `json:"result_list,omitempty" xml:"result_list>benefit_send_result_ext,omitempty"`
    
    
    // uniqueId
    
    UniqueId   string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`

    
}
