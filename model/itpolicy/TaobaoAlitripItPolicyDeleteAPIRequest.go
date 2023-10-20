package itpolicy

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItPolicyDeleteAPIRequest 【国际机票销售规则】单条删除 API请求
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoAlitripItPolicyDeleteAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 接入方产品id
	_outId string
	// 淘宝政策id
	_taobaoId int64
}

// NewTaobaoAlitripItPolicyDeleteRequest 初始化TaobaoAlitripItPolicyDeleteAPIRequest对象
func NewTaobaoAlitripItPolicyDeleteRequest() *TaobaoAlitripItPolicyDeleteAPIRequest {
	return &TaobaoAlitripItPolicyDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripItPolicyDeleteAPIRequest) Reset() {
	r._extendAttributes = ""
	r._outId = ""
	r._taobaoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyDeleteAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 接入方产品id
func (r *TaobaoAlitripItPolicyDeleteAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetOutId() string {
	return r._outId
}

// SetTaobaoId is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyDeleteAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r TaobaoAlitripItPolicyDeleteAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}

var poolTaobaoAlitripItPolicyDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripItPolicyDeleteRequest()
	},
}

// GetTaobaoAlitripItPolicyDeleteRequest 从 sync.Pool 获取 TaobaoAlitripItPolicyDeleteAPIRequest
func GetTaobaoAlitripItPolicyDeleteAPIRequest() *TaobaoAlitripItPolicyDeleteAPIRequest {
	return poolTaobaoAlitripItPolicyDeleteAPIRequest.Get().(*TaobaoAlitripItPolicyDeleteAPIRequest)
}

// ReleaseTaobaoAlitripItPolicyDeleteAPIRequest 将 TaobaoAlitripItPolicyDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripItPolicyDeleteAPIRequest(v *TaobaoAlitripItPolicyDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAlitripItPolicyDeleteAPIRequest.Put(v)
}
