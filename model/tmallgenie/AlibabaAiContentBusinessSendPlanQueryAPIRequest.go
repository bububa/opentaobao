package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessSendPlanQueryAPIRequest 内容商业化发放权益查询 API请求
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
type AlibabaAiContentBusinessSendPlanQueryAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *PurchaseForOuterDto
}

// NewAlibabaAiContentBusinessSendPlanQueryRequest 初始化AlibabaAiContentBusinessSendPlanQueryAPIRequest对象
func NewAlibabaAiContentBusinessSendPlanQueryRequest() *AlibabaAiContentBusinessSendPlanQueryAPIRequest {
	return &AlibabaAiContentBusinessSendPlanQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiContentBusinessSendPlanQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.send.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiContentBusinessSendPlanQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *AlibabaAiContentBusinessSendPlanQueryAPIRequest) SetParam0(_param0 *PurchaseForOuterDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAiContentBusinessSendPlanQueryAPIRequest) GetParam0() *PurchaseForOuterDto {
	return r._param0
}
