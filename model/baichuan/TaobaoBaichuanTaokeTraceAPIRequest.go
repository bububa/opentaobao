package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuantaoketraceAPIRequest 百川淘客打点 API请求
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
type TaobaobaichuantaoketraceAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuantaoketraceRequest 初始化TaobaobaichuantaoketraceAPIRequest对象
func NewTaobaobaichuantaoketraceRequest() *TaobaobaichuantaoketraceAPIRequest {
	return &TaobaobaichuantaoketraceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuantaoketraceAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.taoke.trace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuantaoketraceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuantaoketraceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuantaoketraceAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuantaoketraceAPIRequest) GetName() string {
	return r._name
}
