package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务优惠规则写入 API请求
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入
*/
type AlibabaAsrDataservicePromotionruleWriteRequest struct {
    model.Params
    // 入参对象
    _poskeyPromotionRuleDto   *PosKeyPromotionRuleDto
}

// 初始化AlibabaAsrDataservicePromotionruleWriteRequest对象
func NewAlibabaAsrDataservicePromotionruleWriteRequest() *AlibabaAsrDataservicePromotionruleWriteRequest{
    return &AlibabaAsrDataservicePromotionruleWriteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.write"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PoskeyPromotionRuleDto Setter
// 入参对象
func (r *AlibabaAsrDataservicePromotionruleWriteRequest) SetPoskeyPromotionRuleDto(_poskeyPromotionRuleDto *PosKeyPromotionRuleDto) error {
    r._poskeyPromotionRuleDto = _poskeyPromotionRuleDto
    r.Set("poskey_promotion_rule_dto", _poskeyPromotionRuleDto)
    return nil
}

// PoskeyPromotionRuleDto Getter
func (r AlibabaAsrDataservicePromotionruleWriteRequest) GetPoskeyPromotionRuleDto() *PosKeyPromotionRuleDto {
    return r._poskeyPromotionRuleDto
}
