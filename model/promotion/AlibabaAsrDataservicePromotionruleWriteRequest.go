package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务优惠规则写入 APIRequest
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入
*/
type AlibabaAsrDataservicePromotionruleWriteRequest struct {
    model.Params

    // 入参对象
    poskeyPromotionRuleDto   *PosKeyPromotionRuleDto 

}

func NewAlibabaAsrDataservicePromotionruleWriteRequest() *AlibabaAsrDataservicePromotionruleWriteRequest{
    return &AlibabaAsrDataservicePromotionruleWriteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.write"
}

func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAsrDataservicePromotionruleWriteRequest) SetPoskeyPromotionRuleDto(poskeyPromotionRuleDto *PosKeyPromotionRuleDto) error {
    r.poskeyPromotionRuleDto = poskeyPromotionRuleDto
    r.Set("poskey_promotion_rule_dto", poskeyPromotionRuleDto)
    return nil
}

func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetPoskeyPromotionRuleDto() *PosKeyPromotionRuleDto {
    return r.poskeyPromotionRuleDto
}

