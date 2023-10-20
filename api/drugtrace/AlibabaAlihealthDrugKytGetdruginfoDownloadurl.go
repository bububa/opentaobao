package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetdruginfodownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
func Alibabaalihealthdrugkytgetdruginfodownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetdruginfodownloadurlAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetdruginfodownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
