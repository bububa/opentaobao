package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressordertmscancelAPIRequest 服务商上门取退时间取消接口 API请求
// taobao.logistics.express.order.tms.cancel
//
// 服务商上门取退时间取消接口
type TaobaologisticsexpressordertmscancelAPIRequest struct {
	model.Params
	// tms取消参数
	_tmsToMscCancelOrderRequest *Tms2mscCancelOrderRequest
}

// NewTaobaologisticsexpressordertmscancelRequest 初始化TaobaologisticsexpressordertmscancelAPIRequest对象
func NewTaobaologisticsexpressordertmscancelRequest() *TaobaologisticsexpressordertmscancelAPIRequest {
	return &TaobaologisticsexpressordertmscancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressordertmscancelAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.tms.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressordertmscancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressordertmscancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsToMscCancelOrderRequest is TmsToMscCancelOrderRequest Setter
// tms取消参数
func (r *TaobaologisticsexpressordertmscancelAPIRequest) SetTmsToMscCancelOrderRequest(_tmsToMscCancelOrderRequest *Tms2mscCancelOrderRequest) error {
	r._tmsToMscCancelOrderRequest = _tmsToMscCancelOrderRequest
	r.Set("tms_to_msc_cancel_order_request", _tmsToMscCancelOrderRequest)
	return nil
}

// GetTmsToMscCancelOrderRequest TmsToMscCancelOrderRequest Getter
func (r TaobaologisticsexpressordertmscancelAPIRequest) GetTmsToMscCancelOrderRequest() *Tms2mscCancelOrderRequest {
	return r._tmsToMscCancelOrderRequest
}
