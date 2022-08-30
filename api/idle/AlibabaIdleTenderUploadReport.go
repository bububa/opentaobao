package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderUploadReport 服务商上传验货报告同步给闲鱼
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
func AlibabaIdleTenderUploadReport(clt *core.SDKClient, req *idle.AlibabaIdleTenderUploadReportAPIRequest, session string) (*idle.AlibabaIdleTenderUploadReportAPIResponse, error) {
	var resp idle.AlibabaIdleTenderUploadReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
