package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeErrorReport 码信息错误上报
// alibaba.alihealth.drug.code.error.report
//
// 提供码信息错误上报功能，用于数据校对
func AlibabaAlihealthDrugCodeErrorReport(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeErrorReportAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeErrorReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
