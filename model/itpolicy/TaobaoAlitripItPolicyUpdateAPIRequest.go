package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitpolicyupdateAPIRequest 【国际机票销售规则】单条更新 API请求
// taobao.alitrip.it.policy.update
//
// 销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
type TaobaoalitripitpolicyupdateAPIRequest struct {
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

// NewTaobaoalitripitpolicyupdateRequest 初始化TaobaoalitripitpolicyupdateAPIRequest对象
func NewTaobaoalitripitpolicyupdateRequest() *TaobaoalitripitpolicyupdateAPIRequest {
	return &TaobaoalitripitpolicyupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripitpolicyupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripitpolicyupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripitpolicyupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoalitripitpolicyupdateAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoalitripitpolicyupdateAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 接入方产品id
func (r *TaobaoalitripitpolicyupdateAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoalitripitpolicyupdateAPIRequest) GetOutId() string {
	return r._outId
}

// SetTaobaoId is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoalitripitpolicyupdateAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r TaobaoalitripitpolicyupdateAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}

// SetTopPolicyDo is TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoalitripitpolicyupdateAPIRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
	r._topPolicyDo = _topPolicyDo
	r.Set("top_policy_do", _topPolicyDo)
	return nil
}

// GetTopPolicyDo TopPolicyDo Getter
func (r TaobaoalitripitpolicyupdateAPIRequest) GetTopPolicyDo() *TopPolicyDo {
	return r._topPolicyDo
}
