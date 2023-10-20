package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformAccountQueryInfo 查询账户信息
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
func AlibabaFundplatformAccountQueryInfo(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountQueryInfoAPIRequest, resp *fundplatform.AlibabaFundplatformAccountQueryInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
