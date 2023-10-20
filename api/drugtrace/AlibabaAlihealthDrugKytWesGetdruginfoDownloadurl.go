package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesgetdruginfodownloadurl 药品全量数据下载
// alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl
//
// 药品全量数据下载
func Alibabaalihealthdrugkytwesgetdruginfodownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesgetdruginfodownloadurlAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesgetdruginfodownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesgetdruginfodownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
