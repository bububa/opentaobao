package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageexceptionReportAPIRequest 无主件回告 API请求
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
type TaobaoLogisticsWmsPackageexceptionReportAPIRequest struct {
	model.Params
	// 请求
	_reportExceptionRequest *ReportExceptionRequest
}

// NewTaobaoLogisticsWmsPackageexceptionReportRequest 初始化TaobaoLogisticsWmsPackageexceptionReportAPIRequest对象
func NewTaobaoLogisticsWmsPackageexceptionReportRequest() *TaobaoLogisticsWmsPackageexceptionReportAPIRequest {
	return &TaobaoLogisticsWmsPackageexceptionReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsPackageexceptionReportAPIRequest) Reset() {
	r._reportExceptionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsPackageexceptionReportAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageexception.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsPackageexceptionReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsPackageexceptionReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportExceptionRequest is ReportExceptionRequest Setter
// 请求
func (r *TaobaoLogisticsWmsPackageexceptionReportAPIRequest) SetReportExceptionRequest(_reportExceptionRequest *ReportExceptionRequest) error {
	r._reportExceptionRequest = _reportExceptionRequest
	r.Set("report_exception_request", _reportExceptionRequest)
	return nil
}

// GetReportExceptionRequest ReportExceptionRequest Getter
func (r TaobaoLogisticsWmsPackageexceptionReportAPIRequest) GetReportExceptionRequest() *ReportExceptionRequest {
	return r._reportExceptionRequest
}

var poolTaobaoLogisticsWmsPackageexceptionReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsPackageexceptionReportRequest()
	},
}

// GetTaobaoLogisticsWmsPackageexceptionReportRequest 从 sync.Pool 获取 TaobaoLogisticsWmsPackageexceptionReportAPIRequest
func GetTaobaoLogisticsWmsPackageexceptionReportAPIRequest() *TaobaoLogisticsWmsPackageexceptionReportAPIRequest {
	return poolTaobaoLogisticsWmsPackageexceptionReportAPIRequest.Get().(*TaobaoLogisticsWmsPackageexceptionReportAPIRequest)
}

// ReleaseTaobaoLogisticsWmsPackageexceptionReportAPIRequest 将 TaobaoLogisticsWmsPackageexceptionReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageexceptionReportAPIRequest(v *TaobaoLogisticsWmsPackageexceptionReportAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageexceptionReportAPIRequest.Put(v)
}
