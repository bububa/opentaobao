package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderUploadReportAPIRequest 服务商上传验货报告同步给闲鱼 API请求
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
type AlibabaIdleTenderUploadReportAPIRequest struct {
	model.Params
	// 回收商质检报告
	_param0 *InspectionReport
}

// NewAlibabaIdleTenderUploadReportRequest 初始化AlibabaIdleTenderUploadReportAPIRequest对象
func NewAlibabaIdleTenderUploadReportRequest() *AlibabaIdleTenderUploadReportAPIRequest {
	return &AlibabaIdleTenderUploadReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderUploadReportAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.upload.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderUploadReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderUploadReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 回收商质检报告
func (r *AlibabaIdleTenderUploadReportAPIRequest) SetParam0(_param0 *InspectionReport) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleTenderUploadReportAPIRequest) GetParam0() *InspectionReport {
	return r._param0
}
