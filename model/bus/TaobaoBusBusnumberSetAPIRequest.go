package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusbusnumbersetAPIRequest 商家汽车票车次更新通知接口 API请求
// taobao.bus.busnumber.set
//
// 商家汽车票车次更新后，调用该接口通知平台。
type TaobaobusbusnumbersetAPIRequest struct {
	model.Params
	// 车次更新通知参数
	_pushParam *TopBusNumerPushRq
}

// NewTaobaobusbusnumbersetRequest 初始化TaobaobusbusnumbersetAPIRequest对象
func NewTaobaobusbusnumbersetRequest() *TaobaobusbusnumbersetAPIRequest {
	return &TaobaobusbusnumbersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusbusnumbersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusbusnumbersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusbusnumbersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushParam is PushParam Setter
// 车次更新通知参数
func (r *TaobaobusbusnumbersetAPIRequest) SetPushParam(_pushParam *TopBusNumerPushRq) error {
	r._pushParam = _pushParam
	r.Set("push_param", _pushParam)
	return nil
}

// GetPushParam PushParam Getter
func (r TaobaobusbusnumbersetAPIRequest) GetPushParam() *TopBusNumerPushRq {
	return r._pushParam
}
