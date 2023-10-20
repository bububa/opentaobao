package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleInspectionReportAPIRequest 鉴定报告 API请求
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
type AlibabaIdleRecycleInspectionReportAPIRequest struct {
	model.Params
	// 鉴定报告
	_inspectionReport *InspectionReport
}

// NewAlibabaIdleRecycleInspectionReportRequest 初始化AlibabaIdleRecycleInspectionReportAPIRequest对象
func NewAlibabaIdleRecycleInspectionReportRequest() *AlibabaIdleRecycleInspectionReportAPIRequest {
	return &AlibabaIdleRecycleInspectionReportAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRecycleInspectionReportAPIRequest) Reset() {
	r._inspectionReport = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.inspection.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInspectionReport is InspectionReport Setter
// 鉴定报告
func (r *AlibabaIdleRecycleInspectionReportAPIRequest) SetInspectionReport(_inspectionReport *InspectionReport) error {
	r._inspectionReport = _inspectionReport
	r.Set("inspection_report", _inspectionReport)
	return nil
}

// GetInspectionReport InspectionReport Getter
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetInspectionReport() *InspectionReport {
	return r._inspectionReport
}

var poolAlibabaIdleRecycleInspectionReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRecycleInspectionReportRequest()
	},
}

// GetAlibabaIdleRecycleInspectionReportRequest 从 sync.Pool 获取 AlibabaIdleRecycleInspectionReportAPIRequest
func GetAlibabaIdleRecycleInspectionReportAPIRequest() *AlibabaIdleRecycleInspectionReportAPIRequest {
	return poolAlibabaIdleRecycleInspectionReportAPIRequest.Get().(*AlibabaIdleRecycleInspectionReportAPIRequest)
}

// ReleaseAlibabaIdleRecycleInspectionReportAPIRequest 将 AlibabaIdleRecycleInspectionReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRecycleInspectionReportAPIRequest(v *AlibabaIdleRecycleInspectionReportAPIRequest) {
	v.Reset()
	poolAlibabaIdleRecycleInspectionReportAPIRequest.Put(v)
}
