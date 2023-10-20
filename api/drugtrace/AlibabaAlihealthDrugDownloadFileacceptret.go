package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugdownloadfileacceptret 企业上传回执
// alibaba.alihealth.drug.download.fileacceptret
//
// 拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
func Alibabaalihealthdrugdownloadfileacceptret(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugdownloadfileacceptretAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugdownloadfileacceptretAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugdownloadfileacceptretAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
