package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytquerydruginfo 码查询药品
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
func Alibabaalihealthdrugkytquerydruginfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytquerydruginfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytquerydruginfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytquerydruginfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
