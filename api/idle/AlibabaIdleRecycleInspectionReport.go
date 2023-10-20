package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleInspectionReport 鉴定报告
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
func AlibabaIdleRecycleInspectionReport(clt *core.SDKClient, req *idle.AlibabaIdleRecycleInspectionReportAPIRequest, resp *idle.AlibabaIdleRecycleInspectionReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
