package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportgettargetreportAPIRequest 定向报告 API请求
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
type AlibabascbpadreportgettargetreportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_targetReportOperation *TargetReportOperationDto
}

// NewAlibabascbpadreportgettargetreportRequest 初始化AlibabascbpadreportgettargetreportAPIRequest对象
func NewAlibabascbpadreportgettargetreportRequest() *AlibabascbpadreportgettargetreportAPIRequest {
	return &AlibabascbpadreportgettargetreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportgettargetreportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.target.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportgettargetreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportgettargetreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportgettargetreportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportgettargetreportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetTargetReportOperation is TargetReportOperation Setter
// 请求参数
func (r *AlibabascbpadreportgettargetreportAPIRequest) SetTargetReportOperation(_targetReportOperation *TargetReportOperationDto) error {
	r._targetReportOperation = _targetReportOperation
	r.Set("target_report_operation", _targetReportOperation)
	return nil
}

// GetTargetReportOperation TargetReportOperation Getter
func (r AlibabascbpadreportgettargetreportAPIRequest) GetTargetReportOperation() *TargetReportOperationDto {
	return r._targetReportOperation
}
