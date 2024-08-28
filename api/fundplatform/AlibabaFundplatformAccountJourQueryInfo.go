package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformAccountJourQueryInfo 查询账户流水信息
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
func AlibabaFundplatformAccountJourQueryInfo(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIRequest, resp *fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
