package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanpayresultqueryAPIRequest 百川支付完成回调 API请求
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
type TaobaobaichuanpayresultqueryAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanpayresultqueryRequest 初始化TaobaobaichuanpayresultqueryAPIRequest对象
func NewTaobaobaichuanpayresultqueryRequest() *TaobaobaichuanpayresultqueryAPIRequest {
	return &TaobaobaichuanpayresultqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanpayresultqueryAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.payresult.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanpayresultqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanpayresultqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanpayresultqueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanpayresultqueryAPIRequest) GetName() string {
	return r._name
}
