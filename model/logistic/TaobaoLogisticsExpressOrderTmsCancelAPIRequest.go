package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderTmsCancelAPIRequest 服务商上门取退时间取消接口 API请求
// taobao.logistics.express.order.tms.cancel
//
// 服务商上门取退时间取消接口
type TaobaoLogisticsExpressOrderTmsCancelAPIRequest struct {
	model.Params
	// tms取消参数
	_tmsToMscCancelOrderRequest *Tms2MscCancelOrderRequest
}

// NewTaobaoLogisticsExpressOrderTmsCancelRequest 初始化TaobaoLogisticsExpressOrderTmsCancelAPIRequest对象
func NewTaobaoLogisticsExpressOrderTmsCancelRequest() *TaobaoLogisticsExpressOrderTmsCancelAPIRequest {
	return &TaobaoLogisticsExpressOrderTmsCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressOrderTmsCancelAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.tms.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressOrderTmsCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressOrderTmsCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsToMscCancelOrderRequest is TmsToMscCancelOrderRequest Setter
// tms取消参数
func (r *TaobaoLogisticsExpressOrderTmsCancelAPIRequest) SetTmsToMscCancelOrderRequest(_tmsToMscCancelOrderRequest *Tms2MscCancelOrderRequest) error {
	r._tmsToMscCancelOrderRequest = _tmsToMscCancelOrderRequest
	r.Set("tms_to_msc_cancel_order_request", _tmsToMscCancelOrderRequest)
	return nil
}

// GetTmsToMscCancelOrderRequest TmsToMscCancelOrderRequest Getter
func (r TaobaoLogisticsExpressOrderTmsCancelAPIRequest) GetTmsToMscCancelOrderRequest() *Tms2MscCancelOrderRequest {
	return r._tmsToMscCancelOrderRequest
}
