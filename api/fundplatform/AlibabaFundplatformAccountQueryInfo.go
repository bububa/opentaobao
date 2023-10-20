package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformaccountqueryinfo 查询账户信息
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
func Alibabafundplatformaccountqueryinfo(clt *core.SDKClient, req *fundplatform.AlibabafundplatformaccountqueryinfoAPIRequest, session string) (*fundplatform.AlibabafundplatformaccountqueryinfoAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformaccountqueryinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
