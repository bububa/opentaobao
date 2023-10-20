package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlereportresultupload 服务商上传验货报告
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
func Alibabaidlereportresultupload(clt *core.SDKClient, req *idle.AlibabaidlereportresultuploadAPIRequest, session string) (*idle.AlibabaidlereportresultuploadAPIResponse, error) {
	var resp idle.AlibabaidlereportresultuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
