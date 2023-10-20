package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurl 获取重点追溯品种明细下载URL
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
func Alibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurl(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljggetkeyflagdruginfodownloadurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
