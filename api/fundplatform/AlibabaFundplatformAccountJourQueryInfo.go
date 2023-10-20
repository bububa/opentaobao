package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformAccountJourQueryInfo 查询账户流水信息
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
func AlibabaFundplatformAccountJourQueryInfo(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIRequest, resp *fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
