package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动关联的权益详情获取 APIResponse
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取
*/
type TaobaoPromotionBenefitActivityDetailGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionBenefitActivityDetailGetResponse
}

type TaobaoPromotionBenefitActivityDetailGetResponse struct {
    XMLName xml.Name `xml:"promotion_benefit_activity_detail_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 活动关联的权益详情列表
    
    RelationBenefitDetails   string `json:"relation_benefit_details,omitempty" xml:"relation_benefit_details,omitempty"`

    
}
