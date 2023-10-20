package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmcancelordersetAPIRequest 线下自助机未付款取消订单 API请求
// taobao.bus.tvmcancelorder.set
//
// 自助机汽车票未付款取消订单
type TaobaobustvmcancelordersetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaobustvmcancelordersetRequest 初始化TaobaobustvmcancelordersetAPIRequest对象
func NewTaobaobustvmcancelordersetRequest() *TaobaobustvmcancelordersetAPIRequest {
	return &TaobaobustvmcancelordersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmcancelordersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcancelorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmcancelordersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmcancelordersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobustvmcancelordersetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmcancelordersetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}
