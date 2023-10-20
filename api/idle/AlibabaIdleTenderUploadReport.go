package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderUploadReport 服务商上传验货报告同步给闲鱼
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
func AlibabaIdleTenderUploadReport(clt *core.SDKClient, req *idle.AlibabaIdleTenderUploadReportAPIRequest, resp *idle.AlibabaIdleTenderUploadReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
