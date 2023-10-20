package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticserpdeliverycutAPIRequest ERP发起配拦截 API请求
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
type TaobaologisticserpdeliverycutAPIRequest struct {
	model.Params
	// 请求
	_cutOffDeliveryProcessRequest *CutOffDeliveryProcessRequest
}

// NewTaobaologisticserpdeliverycutRequest 初始化TaobaologisticserpdeliverycutAPIRequest对象
func NewTaobaologisticserpdeliverycutRequest() *TaobaologisticserpdeliverycutAPIRequest {
	return &TaobaologisticserpdeliverycutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticserpdeliverycutAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.erp.delivery.cut"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticserpdeliverycutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticserpdeliverycutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCutOffDeliveryProcessRequest is CutOffDeliveryProcessRequest Setter
// 请求
func (r *TaobaologisticserpdeliverycutAPIRequest) SetCutOffDeliveryProcessRequest(_cutOffDeliveryProcessRequest *CutOffDeliveryProcessRequest) error {
	r._cutOffDeliveryProcessRequest = _cutOffDeliveryProcessRequest
	r.Set("cut_off_delivery_process_request", _cutOffDeliveryProcessRequest)
	return nil
}

// GetCutOffDeliveryProcessRequest CutOffDeliveryProcessRequest Getter
func (r TaobaologisticserpdeliverycutAPIRequest) GetCutOffDeliveryProcessRequest() *CutOffDeliveryProcessRequest {
	return r._cutOffDeliveryProcessRequest
}
