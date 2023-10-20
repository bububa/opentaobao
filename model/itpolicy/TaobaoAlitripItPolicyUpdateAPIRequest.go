package itpolicy

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) Reset() {
	r._extendAttributes = ""
	r._outId = ""
	r._taobaoId = 0
	r._topPolicyDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 接入方产品id
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetOutId() string {
	return r._outId
}

// SetTaobaoId is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}

// SetTopPolicyDo is TopPolicyDo Setter
// 国际机票销售规则
func (r *TaobaoAlitripItPolicyUpdateAPIRequest) SetTopPolicyDo(_topPolicyDo *TopPolicyDo) error {
	r._topPolicyDo = _topPolicyDo
	r.Set("top_policy_do", _topPolicyDo)
	return nil
}

// GetTopPolicyDo TopPolicyDo Getter
func (r TaobaoAlitripItPolicyUpdateAPIRequest) GetTopPolicyDo() *TopPolicyDo {
	return r._topPolicyDo
}

var poolTaobaoAlitripItPolicyUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripItPolicyUpdateRequest()
	},
}

// GetTaobaoAlitripItPolicyUpdateRequest 从 sync.Pool 获取 TaobaoAlitripItPolicyUpdateAPIRequest
func GetTaobaoAlitripItPolicyUpdateAPIRequest() *TaobaoAlitripItPolicyUpdateAPIRequest {
	return poolTaobaoAlitripItPolicyUpdateAPIRequest.Get().(*TaobaoAlitripItPolicyUpdateAPIRequest)
}

// ReleaseTaobaoAlitripItPolicyUpdateAPIRequest 将 TaobaoAlitripItPolicyUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripItPolicyUpdateAPIRequest(v *TaobaoAlitripItPolicyUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripItPolicyUpdateAPIRequest.Put(v)
}
