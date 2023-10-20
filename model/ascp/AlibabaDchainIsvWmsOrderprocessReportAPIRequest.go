package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainIsvWmsOrderprocessReportAPIRequest 仓作业信息同步 API请求
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
type AlibabaDchainIsvWmsOrderprocessReportAPIRequest struct {
	model.Params
	// 请求入参
	_orderProcessReportRequest *WmsOrderProcessReportRequest
}

// NewAlibabaDchainIsvWmsOrderprocessReportRequest 初始化AlibabaDchainIsvWmsOrderprocessReportAPIRequest对象
func NewAlibabaDchainIsvWmsOrderprocessReportRequest() *AlibabaDchainIsvWmsOrderprocessReportAPIRequest {
	return &AlibabaDchainIsvWmsOrderprocessReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainIsvWmsOrderprocessReportAPIRequest) Reset() {
	r._orderProcessReportRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.isv.wms.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderProcessReportRequest is OrderProcessReportRequest Setter
// 请求入参
func (r *AlibabaDchainIsvWmsOrderprocessReportAPIRequest) SetOrderProcessReportRequest(_orderProcessReportRequest *WmsOrderProcessReportRequest) error {
	r._orderProcessReportRequest = _orderProcessReportRequest
	r.Set("order_process_report_request", _orderProcessReportRequest)
	return nil
}

// GetOrderProcessReportRequest OrderProcessReportRequest Getter
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetOrderProcessReportRequest() *WmsOrderProcessReportRequest {
	return r._orderProcessReportRequest
}

var poolAlibabaDchainIsvWmsOrderprocessReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainIsvWmsOrderprocessReportRequest()
	},
}

// GetAlibabaDchainIsvWmsOrderprocessReportRequest 从 sync.Pool 获取 AlibabaDchainIsvWmsOrderprocessReportAPIRequest
func GetAlibabaDchainIsvWmsOrderprocessReportAPIRequest() *AlibabaDchainIsvWmsOrderprocessReportAPIRequest {
	return poolAlibabaDchainIsvWmsOrderprocessReportAPIRequest.Get().(*AlibabaDchainIsvWmsOrderprocessReportAPIRequest)
}

// ReleaseAlibabaDchainIsvWmsOrderprocessReportAPIRequest 将 AlibabaDchainIsvWmsOrderprocessReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainIsvWmsOrderprocessReportAPIRequest(v *AlibabaDchainIsvWmsOrderprocessReportAPIRequest) {
	v.Reset()
	poolAlibabaDchainIsvWmsOrderprocessReportAPIRequest.Put(v)
}
