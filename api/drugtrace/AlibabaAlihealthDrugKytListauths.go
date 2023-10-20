package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytlistauths 企业搜索自己授权的物流企业
// alibaba.alihealth.drug.kyt.listauths
//
// 企业搜索自己授权的物流企业
func Alibabaalihealthdrugkytlistauths(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytlistauthsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytlistauthsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytlistauthsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
