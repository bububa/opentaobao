package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusBusnumberSetAPIRequest 商家汽车票车次更新通知接口 API请求
// taobao.bus.busnumber.set
//
// 商家汽车票车次更新后，调用该接口通知平台。
type TaobaoBusBusnumberSetAPIRequest struct {
	model.Params
	// 车次更新通知参数
	_pushParam *TopBusNumerPushRq
}

// NewTaobaoBusBusnumberSetRequest 初始化TaobaoBusBusnumberSetAPIRequest对象
func NewTaobaoBusBusnumberSetRequest() *TaobaoBusBusnumberSetAPIRequest {
	return &TaobaoBusBusnumberSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushParam Setter
// 车次更新通知参数
func (r *TaobaoBusBusnumberSetAPIRequest) SetPushParam(_pushParam *TopBusNumerPushRq) error {
	r._pushParam = _pushParam
	r.Set("push_param", _pushParam)
	return nil
}

// Get PushParam Getter
func (r TaobaoBusBusnumberSetAPIRequest) GetPushParam() *TopBusNumerPushRq {
	return r._pushParam
}
