package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanPayresultQueryAPIRequest 百川支付完成回调 API请求
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
type TaobaoBaichuanPayresultQueryAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanPayresultQueryRequest 初始化TaobaoBaichuanPayresultQueryAPIRequest对象
func NewTaobaoBaichuanPayresultQueryRequest() *TaobaoBaichuanPayresultQueryAPIRequest {
	return &TaobaoBaichuanPayresultQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.payresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// name
func (r *TaobaoBaichuanPayresultQueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoBaichuanPayresultQueryAPIRequest) GetName() string {
	return r._name
}
