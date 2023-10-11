package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest 仓作业信息批量同步 API请求
// alibaba.dchain.isv.wms.orderprocess.batch.report
//
// 仓作业信息批量同步
type AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest struct {
	model.Params
	// 入参
	_orderProcessBatchReportRequest *WmsOrderProcessBatchReportRequest
}

// NewAlibabaDchainIsvWmsOrderprocessBatchReportRequest 初始化AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest对象
func NewAlibabaDchainIsvWmsOrderprocessBatchReportRequest() *AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest {
	return &AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.isv.wms.orderprocess.batch.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderProcessBatchReportRequest is OrderProcessBatchReportRequest Setter
// 入参
func (r *AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest) SetOrderProcessBatchReportRequest(_orderProcessBatchReportRequest *WmsOrderProcessBatchReportRequest) error {
	r._orderProcessBatchReportRequest = _orderProcessBatchReportRequest
	r.Set("order_process_batch_report_request", _orderProcessBatchReportRequest)
	return nil
}

// GetOrderProcessBatchReportRequest OrderProcessBatchReportRequest Getter
func (r AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest) GetOrderProcessBatchReportRequest() *WmsOrderProcessBatchReportRequest {
	return r._orderProcessBatchReportRequest
}
