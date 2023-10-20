package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsErpDeliveryCutAPIRequest ERP发起配拦截 API请求
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
type TaobaoLogisticsErpDeliveryCutAPIRequest struct {
	model.Params
	// 请求
	_cutOffDeliveryProcessRequest *CutOffDeliveryProcessRequest
}

// NewTaobaoLogisticsErpDeliveryCutRequest 初始化TaobaoLogisticsErpDeliveryCutAPIRequest对象
func NewTaobaoLogisticsErpDeliveryCutRequest() *TaobaoLogisticsErpDeliveryCutAPIRequest {
	return &TaobaoLogisticsErpDeliveryCutAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsErpDeliveryCutAPIRequest) Reset() {
	r._cutOffDeliveryProcessRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsErpDeliveryCutAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.erp.delivery.cut"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsErpDeliveryCutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsErpDeliveryCutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCutOffDeliveryProcessRequest is CutOffDeliveryProcessRequest Setter
// 请求
func (r *TaobaoLogisticsErpDeliveryCutAPIRequest) SetCutOffDeliveryProcessRequest(_cutOffDeliveryProcessRequest *CutOffDeliveryProcessRequest) error {
	r._cutOffDeliveryProcessRequest = _cutOffDeliveryProcessRequest
	r.Set("cut_off_delivery_process_request", _cutOffDeliveryProcessRequest)
	return nil
}

// GetCutOffDeliveryProcessRequest CutOffDeliveryProcessRequest Getter
func (r TaobaoLogisticsErpDeliveryCutAPIRequest) GetCutOffDeliveryProcessRequest() *CutOffDeliveryProcessRequest {
	return r._cutOffDeliveryProcessRequest
}

var poolTaobaoLogisticsErpDeliveryCutAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsErpDeliveryCutRequest()
	},
}

// GetTaobaoLogisticsErpDeliveryCutRequest 从 sync.Pool 获取 TaobaoLogisticsErpDeliveryCutAPIRequest
func GetTaobaoLogisticsErpDeliveryCutAPIRequest() *TaobaoLogisticsErpDeliveryCutAPIRequest {
	return poolTaobaoLogisticsErpDeliveryCutAPIRequest.Get().(*TaobaoLogisticsErpDeliveryCutAPIRequest)
}

// ReleaseTaobaoLogisticsErpDeliveryCutAPIRequest 将 TaobaoLogisticsErpDeliveryCutAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsErpDeliveryCutAPIRequest(v *TaobaoLogisticsErpDeliveryCutAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsErpDeliveryCutAPIRequest.Put(v)
}
