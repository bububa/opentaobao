package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetTargetReportAPIRequest 定向报告 API请求
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
type AlibabaScbpAdReportGetTargetReportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_targetReportOperation *TargetReportOperationDto
}

// NewAlibabaScbpAdReportGetTargetReportRequest 初始化AlibabaScbpAdReportGetTargetReportAPIRequest对象
func NewAlibabaScbpAdReportGetTargetReportRequest() *AlibabaScbpAdReportGetTargetReportAPIRequest {
	return &AlibabaScbpAdReportGetTargetReportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportGetTargetReportAPIRequest) Reset() {
	r._topContext = nil
	r._targetReportOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.target.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetTargetReportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetTargetReportOperation is TargetReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetTargetReportAPIRequest) SetTargetReportOperation(_targetReportOperation *TargetReportOperationDto) error {
	r._targetReportOperation = _targetReportOperation
	r.Set("target_report_operation", _targetReportOperation)
	return nil
}

// GetTargetReportOperation TargetReportOperation Getter
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetTargetReportOperation() *TargetReportOperationDto {
	return r._targetReportOperation
}

var poolAlibabaScbpAdReportGetTargetReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportGetTargetReportRequest()
	},
}

// GetAlibabaScbpAdReportGetTargetReportRequest 从 sync.Pool 获取 AlibabaScbpAdReportGetTargetReportAPIRequest
func GetAlibabaScbpAdReportGetTargetReportAPIRequest() *AlibabaScbpAdReportGetTargetReportAPIRequest {
	return poolAlibabaScbpAdReportGetTargetReportAPIRequest.Get().(*AlibabaScbpAdReportGetTargetReportAPIRequest)
}

// ReleaseAlibabaScbpAdReportGetTargetReportAPIRequest 将 AlibabaScbpAdReportGetTargetReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportGetTargetReportAPIRequest(v *AlibabaScbpAdReportGetTargetReportAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportGetTargetReportAPIRequest.Put(v)
}
