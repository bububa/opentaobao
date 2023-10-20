package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleinspectionreportAPIRequest 鉴定报告 API请求
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
type AlibabaidlerecycleinspectionreportAPIRequest struct {
	model.Params
	// 鉴定报告
	_inspectionReport *InspectionReport
}

// NewAlibabaidlerecycleinspectionreportRequest 初始化AlibabaidlerecycleinspectionreportAPIRequest对象
func NewAlibabaidlerecycleinspectionreportRequest() *AlibabaidlerecycleinspectionreportAPIRequest {
	return &AlibabaidlerecycleinspectionreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleinspectionreportAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.inspection.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleinspectionreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleinspectionreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInspectionReport is InspectionReport Setter
// 鉴定报告
func (r *AlibabaidlerecycleinspectionreportAPIRequest) SetInspectionReport(_inspectionReport *InspectionReport) error {
	r._inspectionReport = _inspectionReport
	r.Set("inspection_report", _inspectionReport)
	return nil
}

// GetInspectionReport InspectionReport Getter
func (r AlibabaidlerecycleinspectionreportAPIRequest) GetInspectionReport() *InspectionReport {
	return r._inspectionReport
}
