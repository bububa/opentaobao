package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitpolicydeleteAPIRequest 【国际机票销售规则】单条删除 API请求
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoalitripitpolicydeleteAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 接入方产品id
	_outId string
	// 淘宝政策id
	_taobaoId int64
}

// NewTaobaoalitripitpolicydeleteRequest 初始化TaobaoalitripitpolicydeleteAPIRequest对象
func NewTaobaoalitripitpolicydeleteRequest() *TaobaoalitripitpolicydeleteAPIRequest {
	return &TaobaoalitripitpolicydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripitpolicydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripitpolicydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripitpolicydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 扩展字段
func (r *TaobaoalitripitpolicydeleteAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoalitripitpolicydeleteAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 接入方产品id
func (r *TaobaoalitripitpolicydeleteAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoalitripitpolicydeleteAPIRequest) GetOutId() string {
	return r._outId
}

// SetTaobaoId is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoalitripitpolicydeleteAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r TaobaoalitripitpolicydeleteAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}
