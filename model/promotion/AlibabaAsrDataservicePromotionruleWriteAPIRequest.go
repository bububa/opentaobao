package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaasrdataservicepromotionrulewriteAPIRequest 业务优惠规则写入 API请求
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
type AlibabaasrdataservicepromotionrulewriteAPIRequest struct {
	model.Params
	// 入参对象
	_poskeyPromotionRuleDto *PosKeyPromotionRuleDto
}

// NewAlibabaasrdataservicepromotionrulewriteRequest 初始化AlibabaasrdataservicepromotionrulewriteAPIRequest对象
func NewAlibabaasrdataservicepromotionrulewriteRequest() *AlibabaasrdataservicepromotionrulewriteAPIRequest {
	return &AlibabaasrdataservicepromotionrulewriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaasrdataservicepromotionrulewriteAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaasrdataservicepromotionrulewriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaasrdataservicepromotionrulewriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoskeyPromotionRuleDto is PoskeyPromotionRuleDto Setter
// 入参对象
func (r *AlibabaasrdataservicepromotionrulewriteAPIRequest) SetPoskeyPromotionRuleDto(_poskeyPromotionRuleDto *PosKeyPromotionRuleDto) error {
	r._poskeyPromotionRuleDto = _poskeyPromotionRuleDto
	r.Set("poskey_promotion_rule_dto", _poskeyPromotionRuleDto)
	return nil
}

// GetPoskeyPromotionRuleDto PoskeyPromotionRuleDto Getter
func (r AlibabaasrdataservicepromotionrulewriteAPIRequest) GetPoskeyPromotionRuleDto() *PosKeyPromotionRuleDto {
	return r._poskeyPromotionRuleDto
}
