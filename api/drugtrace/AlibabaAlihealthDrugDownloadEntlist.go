package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugdownloadentlist 企业下载列表
// alibaba.alihealth.drug.download.entlist
//
// 获取企业的下载文件列表
func Alibabaalihealthdrugdownloadentlist(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugdownloadentlistAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugdownloadentlistAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugdownloadentlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
