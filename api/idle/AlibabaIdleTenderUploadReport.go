package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderuploadreport 服务商上传验货报告同步给闲鱼
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
func Alibabaidletenderuploadreport(clt *core.SDKClient, req *idle.AlibabaidletenderuploadreportAPIRequest, session string) (*idle.AlibabaidletenderuploadreportAPIResponse, error) {
	var resp idle.AlibabaidletenderuploadreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
