package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanTaokeTraceAPIRequest 百川淘客打点 API请求
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
type TaobaoBaichuanTaokeTraceAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanTaokeTraceRequest 初始化TaobaoBaichuanTaokeTraceAPIRequest对象
func NewTaobaoBaichuanTaokeTraceRequest() *TaobaoBaichuanTaokeTraceAPIRequest {
	return &TaobaoBaichuanTaokeTraceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanTaokeTraceAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.taoke.trace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanTaokeTraceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaichuanTaokeTraceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanTaokeTraceAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanTaokeTraceAPIRequest) GetName() string {
	return r._name
}
