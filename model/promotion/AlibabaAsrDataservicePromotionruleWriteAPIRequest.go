package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAsrDataservicePromotionruleWriteAPIRequest
业务优惠规则写入 API请求
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入 */
type AlibabaAsrDataservicePromotionruleWriteAPIRequest struct {
	model.Params
	// 入参对象
	_poskeyPromotionRuleDto *PosKeyPromotionRuleDto
}

// NewAlibabaAsrDataservicePromotionruleWriteRequest 初始化AlibabaAsrDataservicePromotionruleWriteAPIRequest对象
func NewAlibabaAsrDataservicePromotionruleWriteRequest() *AlibabaAsrDataservicePromotionruleWriteAPIRequest {
	return &AlibabaAsrDataservicePromotionruleWriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PoskeyPromotionRuleDto Setter
// 入参对象
func (r *AlibabaAsrDataservicePromotionruleWriteAPIRequest) SetPoskeyPromotionRuleDto(_poskeyPromotionRuleDto *PosKeyPromotionRuleDto) error {
	r._poskeyPromotionRuleDto = _poskeyPromotionRuleDto
	r.Set("poskey_promotion_rule_dto", _poskeyPromotionRuleDto)
	return nil
}

// Get PoskeyPromotionRuleDto Getter
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetPoskeyPromotionRuleDto() *PosKeyPromotionRuleDto {
	return r._poskeyPromotionRuleDto
}
