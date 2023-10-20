package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackageexceptionreportAPIRequest 无主件回告 API请求
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
type TaobaologisticswmspackageexceptionreportAPIRequest struct {
	model.Params
	// 请求
	_reportExceptionRequest *ReportExceptionRequest
}

// NewTaobaologisticswmspackageexceptionreportRequest 初始化TaobaologisticswmspackageexceptionreportAPIRequest对象
func NewTaobaologisticswmspackageexceptionreportRequest() *TaobaologisticswmspackageexceptionreportAPIRequest {
	return &TaobaologisticswmspackageexceptionreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmspackageexceptionreportAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageexception.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmspackageexceptionreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmspackageexceptionreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportExceptionRequest is ReportExceptionRequest Setter
// 请求
func (r *TaobaologisticswmspackageexceptionreportAPIRequest) SetReportExceptionRequest(_reportExceptionRequest *ReportExceptionRequest) error {
	r._reportExceptionRequest = _reportExceptionRequest
	r.Set("report_exception_request", _reportExceptionRequest)
	return nil
}

// GetReportExceptionRequest ReportExceptionRequest Getter
func (r TaobaologisticswmspackageexceptionreportAPIRequest) GetReportExceptionRequest() *ReportExceptionRequest {
	return r._reportExceptionRequest
}
