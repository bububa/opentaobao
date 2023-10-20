package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanorderurlgetAPIRequest 百川订单详情 API请求
// taobao.baichuan.orderurl.get
//
// 百川订单详情
type TaobaobaichuanorderurlgetAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaobaichuanorderurlgetRequest 初始化TaobaobaichuanorderurlgetAPIRequest对象
func NewTaobaobaichuanorderurlgetRequest() *TaobaobaichuanorderurlgetAPIRequest {
	return &TaobaobaichuanorderurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaichuanorderurlgetAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.orderurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaichuanorderurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaichuanorderurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// name
func (r *TaobaobaichuanorderurlgetAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaobaichuanorderurlgetAPIRequest) GetName() string {
	return r._name
}
