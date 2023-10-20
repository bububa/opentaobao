package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleReportResultUploadAPIRequest 服务商上传验货报告 API请求
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
type AlibabaIdleReportResultUploadAPIRequest struct {
	model.Params
	// 参数
	_reportUploadTopCmd *ReportUploadTopCmd
}

// NewAlibabaIdleReportResultUploadRequest 初始化AlibabaIdleReportResultUploadAPIRequest对象
func NewAlibabaIdleReportResultUploadRequest() *AlibabaIdleReportResultUploadAPIRequest {
	return &AlibabaIdleReportResultUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleReportResultUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.report.result.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleReportResultUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleReportResultUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportUploadTopCmd is ReportUploadTopCmd Setter
// 参数
func (r *AlibabaIdleReportResultUploadAPIRequest) SetReportUploadTopCmd(_reportUploadTopCmd *ReportUploadTopCmd) error {
	r._reportUploadTopCmd = _reportUploadTopCmd
	r.Set("report_upload_top_cmd", _reportUploadTopCmd)
	return nil
}

// GetReportUploadTopCmd ReportUploadTopCmd Getter
func (r AlibabaIdleReportResultUploadAPIRequest) GetReportUploadTopCmd() *ReportUploadTopCmd {
	return r._reportUploadTopCmd
}
