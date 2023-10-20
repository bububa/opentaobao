package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsordercancelAPIRequest 回传发货单取消通知 API请求
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
type AlibabadchainaoxiangwmsordercancelAPIRequest struct {
	model.Params
	// 单据取消回传上报请求
	_orderCancelReportRequest *OrderCancelReportRequest
}

// NewAlibabadchainaoxiangwmsordercancelRequest 初始化AlibabadchainaoxiangwmsordercancelAPIRequest对象
func NewAlibabadchainaoxiangwmsordercancelRequest() *AlibabadchainaoxiangwmsordercancelAPIRequest {
	return &AlibabadchainaoxiangwmsordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangwmsordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangwmsordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangwmsordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCancelReportRequest is OrderCancelReportRequest Setter
// 单据取消回传上报请求
func (r *AlibabadchainaoxiangwmsordercancelAPIRequest) SetOrderCancelReportRequest(_orderCancelReportRequest *OrderCancelReportRequest) error {
	r._orderCancelReportRequest = _orderCancelReportRequest
	r.Set("order_cancel_report_request", _orderCancelReportRequest)
	return nil
}

// GetOrderCancelReportRequest OrderCancelReportRequest Getter
func (r AlibabadchainaoxiangwmsordercancelAPIRequest) GetOrderCancelReportRequest() *OrderCancelReportRequest {
	return r._orderCancelReportRequest
}
