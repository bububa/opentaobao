package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformaccountjourqueryinfo 查询账户流水信息
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
func Alibabafundplatformaccountjourqueryinfo(clt *core.SDKClient, req *fundplatform.AlibabafundplatformaccountjourqueryinfoAPIRequest, session string) (*fundplatform.AlibabafundplatformaccountjourqueryinfoAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformaccountjourqueryinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
