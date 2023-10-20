package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusdisableqrcodesetAPIRequest 自助机失效二维码 API请求
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
type TaobaobusdisableqrcodesetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaobusdisableqrcodesetRequest 初始化TaobaobusdisableqrcodesetAPIRequest对象
func NewTaobaobusdisableqrcodesetRequest() *TaobaobusdisableqrcodesetAPIRequest {
	return &TaobaobusdisableqrcodesetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusdisableqrcodesetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.disableqrcode.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusdisableqrcodesetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusdisableqrcodesetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobusdisableqrcodesetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobusdisableqrcodesetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}
