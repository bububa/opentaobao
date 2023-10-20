package tblogistics

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
