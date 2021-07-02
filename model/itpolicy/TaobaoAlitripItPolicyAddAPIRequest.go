package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyAddAPIRequest 【国际机票销售规则】单条新增 API请求
// taobao.alitrip.it.policy.add
//
// 销售规则新增，成功返回taobaoId
type TaobaoAlitripItPolicyAddAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 国际机票销售规则
	_topPolicyDo *TopPolicyDo
}

// NewTaobaoAlitripItPolicyAddRequest 初始化TaobaoAlitripItPolicyAddAPIRequest对象
func NewTaobaoAlitripItPolicyAddRequest() *TaobaoAlitripItPolicyAddAPIRequest {
	return &TaobaoAlitripItPolicyAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyAddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyAddAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// Get ExtendAttributes Getter
func (r TaobaoAlitripItPolicyAddAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// Set is TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoAlitripItPolicyAddAPIRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
	r._topPolicyDo = _topPolicyDo
	r.Set("top_policy_do", _topPolicyDo)
	return nil
}

// Get TopPolicyDo Getter
func (r TaobaoAlitripItPolicyAddAPIRequest) GetTopPolicyDo() *TopPolicyDo {
	return r._topPolicyDo
}
