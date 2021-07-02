package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyUpdateAPIRequest 【国际机票销售规则】单条更新 API请求
// taobao.alitrip.it.policy.update
//
// 销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
type TaobaoAlitripItPolicyUpdateAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 接入方产品id
	_outId string
	// 淘宝政策id
	_taobaoId int64
	// 国际机票销售规则
	_topPolicyDo *TopPolicyDo
}

// NewTaobaoAlitripItPolicyUpdateRequest 初始化TaobaoAlitripItPolicyUpdateAPIRequest对象
func NewTaobaoAlitripItPolicyUpdateRequest() *TaobaoAlitripItPolicyUpdateAPIRequest {
	return &TaobaoAlitripItPolicyUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// Get ExtendAttributes Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// Set is OutId Setter
// 接入方产品id
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// Get OutId Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetOutId() string {
	return r._outId
}

// Set is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// Get TaobaoId Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}

// Set is TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
	r._topPolicyDo = _topPolicyDo
	r.Set("top_policy_do", _topPolicyDo)
	return nil
}

// Get TopPolicyDo Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetTopPolicyDo() *TopPolicyDo {
	return r._topPolicyDo
}
