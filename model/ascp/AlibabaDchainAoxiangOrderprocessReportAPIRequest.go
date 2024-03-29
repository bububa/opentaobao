package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangOrderprocessReportAPIRequest 回传仓内作业节点 API请求
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
type AlibabaDchainAoxiangOrderprocessReportAPIRequest struct {
	model.Params
	// 单据作业状态回传参数
	_orderprocessReportRequest *OrderProcessReportRequest
}

// NewAlibabaDchainAoxiangOrderprocessReportRequest 初始化AlibabaDchainAoxiangOrderprocessReportAPIRequest对象
func NewAlibabaDchainAoxiangOrderprocessReportRequest() *AlibabaDchainAoxiangOrderprocessReportAPIRequest {
	return &AlibabaDchainAoxiangOrderprocessReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangOrderprocessReportAPIRequest) Reset() {
	r._orderprocessReportRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderprocessReportRequest is OrderprocessReportRequest Setter
// 单据作业状态回传参数
func (r *AlibabaDchainAoxiangOrderprocessReportAPIRequest) SetOrderprocessReportRequest(_orderprocessReportRequest *OrderProcessReportRequest) error {
	r._orderprocessReportRequest = _orderprocessReportRequest
	r.Set("orderprocess_report_request", _orderprocessReportRequest)
	return nil
}

// GetOrderprocessReportRequest OrderprocessReportRequest Getter
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetOrderprocessReportRequest() *OrderProcessReportRequest {
	return r._orderprocessReportRequest
}

var poolAlibabaDchainAoxiangOrderprocessReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangOrderprocessReportRequest()
	},
}

// GetAlibabaDchainAoxiangOrderprocessReportRequest 从 sync.Pool 获取 AlibabaDchainAoxiangOrderprocessReportAPIRequest
func GetAlibabaDchainAoxiangOrderprocessReportAPIRequest() *AlibabaDchainAoxiangOrderprocessReportAPIRequest {
	return poolAlibabaDchainAoxiangOrderprocessReportAPIRequest.Get().(*AlibabaDchainAoxiangOrderprocessReportAPIRequest)
}

// ReleaseAlibabaDchainAoxiangOrderprocessReportAPIRequest 将 AlibabaDchainAoxiangOrderprocessReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangOrderprocessReportAPIRequest(v *AlibabaDchainAoxiangOrderprocessReportAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangOrderprocessReportAPIRequest.Put(v)
}
