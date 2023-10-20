package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleReportResultUpload 服务商上传验货报告
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
func AlibabaIdleReportResultUpload(clt *core.SDKClient, req *idle.AlibabaIdleReportResultUploadAPIRequest, resp *idle.AlibabaIdleReportResultUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
