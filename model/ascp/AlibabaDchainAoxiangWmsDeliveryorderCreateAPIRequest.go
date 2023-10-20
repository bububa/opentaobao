package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest 回传仓库接单通知 API请求
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
type AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest struct {
	model.Params
	// 接单回传上报请求
	_deliveryOrderReportRequest *DeliveryOrderReportRequest
}

// NewAlibabadchainaoxiangwmsdeliveryordercreateRequest 初始化AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest对象
func NewAlibabadchainaoxiangwmsdeliveryordercreateRequest() *AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest {
	return &AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.deliveryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrderReportRequest is DeliveryOrderReportRequest Setter
// 接单回传上报请求
func (r *AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest) SetDeliveryOrderReportRequest(_deliveryOrderReportRequest *DeliveryOrderReportRequest) error {
	r._deliveryOrderReportRequest = _deliveryOrderReportRequest
	r.Set("delivery_order_report_request", _deliveryOrderReportRequest)
	return nil
}

// GetDeliveryOrderReportRequest DeliveryOrderReportRequest Getter
func (r AlibabadchainaoxiangwmsdeliveryordercreateAPIRequest) GetDeliveryOrderReportRequest() *DeliveryOrderReportRequest {
	return r._deliveryOrderReportRequest
}
