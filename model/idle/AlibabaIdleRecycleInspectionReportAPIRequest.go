package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleInspectionReportAPIRequest
鉴定报告 API请求
alibaba.idle.recycle.inspection.report

回收商鉴定报告 */
type AlibabaIdleRecycleInspectionReportAPIRequest struct {
	model.Params
	// 鉴定报告
	_inspectionReport *InspectionReport
}

// New
