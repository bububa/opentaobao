package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytfiledownload 处理失败单据下载
// alibaba.alihealth.drug.kyt.filedownload
//
// 处理失败单据下载
func Alibabaalihealthdrugkytfiledownload(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytfiledownloadAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytfiledownloadAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytfiledownloadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
