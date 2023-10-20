package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurl 获取重点追溯品种明细下载URL
// alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
func Alibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydgetkeyflagdruginfodownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
