package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCustbaseGet 客户账户报表基础数据对象
// taobao.simba.rpt.custbase.get
//
// 客户账户报表基础数据对象
func TaobaoSimbaRptCustbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptCustbaseGetAPIRequest, resp *simba.TaobaoSimbaRptCustbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
