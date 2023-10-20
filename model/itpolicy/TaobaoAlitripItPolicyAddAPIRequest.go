package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitpolicyaddAPIRequest 【国际机票销售规则】单条新增 API请求
// taobao.alitrip.it.policy.add
//
// 销售规则新增，成功返回taobaoId
type TaobaoalitripitpolicyaddAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 国际机票销售规则
	_topPolicyDo *TopPolicyDo
}

// NewTaobaoalitripitpolicyaddRequest 初始化TaobaoalitripitpolicyaddAPIRequest对象
func NewTaobaoalitripitpolicyaddRequest() *TaobaoalitripitpolicyaddAPIRequest {
	return &TaobaoalitripitpolicyaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripitpolicyaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripitpolicyaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripitpolicyaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoalitripitpolicyaddAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoalitripitpolicyaddAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetTopPolicyDo is TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoalitripitpolicyaddAPIRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
	r._topPolicyDo = _topPolicyDo
	r.Set("top_policy_do", _topPolicyDo)
	return nil
}

// GetTopPolicyDo TopPolicyDo Getter
func (r TaobaoalitripitpolicyaddAPIRequest) GetTopPolicyDo() *TopPolicyDo {
	return r._topPolicyDo
}
