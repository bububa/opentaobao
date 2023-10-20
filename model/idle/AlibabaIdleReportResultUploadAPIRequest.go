package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlereportresultuploadAPIRequest 服务商上传验货报告 API请求
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
type AlibabaidlereportresultuploadAPIRequest struct {
	model.Params
	// 参数
	_reportUploadTopCmd *ReportUploadTopCmd
}

// NewAlibabaidlereportresultuploadRequest 初始化AlibabaidlereportresultuploadAPIRequest对象
func NewAlibabaidlereportresultuploadRequest() *AlibabaidlereportresultuploadAPIRequest {
	return &AlibabaidlereportresultuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlereportresultuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.report.result.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlereportresultuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlereportresultuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportUploadTopCmd is ReportUploadTopCmd Setter
// 参数
func (r *AlibabaidlereportresultuploadAPIRequest) SetReportUploadTopCmd(_reportUploadTopCmd *ReportUploadTopCmd) error {
	r._reportUploadTopCmd = _reportUploadTopCmd
	r.Set("report_upload_top_cmd", _reportUploadTopCmd)
	return nil
}

// GetReportUploadTopCmd ReportUploadTopCmd Getter
func (r AlibabaidlereportresultuploadAPIRequest) GetReportUploadTopCmd() *ReportUploadTopCmd {
	return r._reportUploadTopCmd
}
